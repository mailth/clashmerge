package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"gopkg.in/yaml.v2"
)

var configDir string

func Init() error {
	// 设置日志输出前缀
	configDir = os.Getenv("CONFIG_DIR")
	if configDir == "" {
		configDir = "/data"
	}
	initLog(filepath.Join(configDir, "/log/errors.log"))
	return nil
}

func main() {
	err := Init()
	if err != nil {
		log.Fatalf("获取运行路径失败")
		return
	}

	// 获取参数
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	log.Infof("服务启动，端口：" + port)
	http.HandleFunc("/", handle)
	err = http.ListenAndServe("0.0.0.0:"+port, nil)
	if err != nil {
		log.Errorf("start server fail, %v", err)
	}
}


func handle(w http.ResponseWriter, r *http.Request) {
	res, err := func() ([]byte, error) {
		configFileName := r.URL.Query().Get("name")
		if configFileName == "" {
			return []byte("{}"), nil
		}
		configItemArr, err := parseConfig(configFileName)
		if err != nil {
			return nil, err
		}
		items, err := processConfig(configItemArr)
		if err != nil {
			return nil, err
		}
		data := map[string]any{}
		res, err := doMerge(data, items)
		if err != nil {
			return nil, err
		}
		return yaml.Marshal(res)
	}()
	if err != nil {
		log.Errorf(err.Error())
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(res)
}
func parseConfig(configFileName string) ([]*ConfigItem, error) {
	// 根据配置名称读取文件
	configFilePath := filepath.Join(configDir, configFileName)
	configContent, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("无法读取配置文件: %v", err)
	}
	log.Infof("成功读取配置文件: %s", configFileName)
	var configItemArr []*ConfigItem
	err = yaml.Unmarshal(configContent, &configItemArr)
	if err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %v", err)
	}
	return configItemArr, nil
}

func processConfig(configItemArr []*ConfigItem) ([]*ConfigItem, error) {
	for i, data:= range configItemArr {
		item:= data
		_type := item.Type
		if _type == "" {
			return nil, errors.New("type is required")
		}
		if _type == "url" {
			_data, err := http.Get(item.URL)
			if err != nil {
				return nil, fmt.Errorf("获取URL失败,url %s, %v", item.URL, err)
			}
			defer _data.Body.Close()
			dataBytes, err := io.ReadAll(_data.Body)
			if err != nil {
				return nil, fmt.Errorf("读取URL内容失败,url %s, %v", item.URL, err)
			}
			var dataMap map[string]any
			err = yaml.Unmarshal(dataBytes, &dataMap)
			if err != nil {
				return nil, fmt.Errorf("解析URL内容失败,url %s, %v", item.URL, err)
			}
			item.Data = dataMap
			configItemArr[i] = item
		}
	}
	return configItemArr, nil
}

func initLog(_filepath string) {
	level := logrus.InfoLevel
	envLevel:= os.Getenv("LOG_LEVEL")
	v, err := logrus.ParseLevel(envLevel)
	if err == nil {
		level = v
	}
	logrus.SetLevel(level)

	lumberjackLogger := &lumberjack.Logger{
			// Log file abbsolute path, os agnostic
			Filename:   filepath.ToSlash(_filepath),
			MaxSize:    10, // MB
			MaxBackups: 5,
		MaxAge:     30, // days
		//Compress:   true, // disabled by default
	}
	// Fork writing into two outputs
	multiWriter := io.MultiWriter(os.Stdout, lumberjackLogger)
	logrus.SetOutput(multiWriter)
}