package main

import (
	"net/http"
	"os"
	"strings"

	"clashmerge/handlers"
	"clashmerge/models"
	"clashmerge/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func main() {
	model, err := models.NewDB()
	if err != nil {
		log.Fatalf("初始化失败: %v", err)
		return
	}

	// 获取参数
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	// 设置 Gin 模式
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// 设置路由
	setupRoutes(r, model)

	log.Infof("server started, port: " + port)
	err = r.Run("0.0.0.0:" + port)
	if err != nil {
		log.Errorf("start server fail, %v", err)
	}
}

// 设置路由
func setupRoutes(r *gin.Engine, model *models.Model) {
	mergeService := service.NewMergeService(model)

	mergeHandler := handlers.NewMergeHandler(model, mergeService)

	r.Any("/", func(c *gin.Context) {
		// 如果有查询参数，则执行配置合并
		logrus.Infof("handle merge, rawQuery: %s", c.Request.URL.RawQuery)
		if c.Request.URL.RawQuery != "" {
			mergeHandler.HandleMerge(c)
			return
		}
		// 否则重定向到管理页面
		c.Writer.WriteHeader(http.StatusFound)
	})

	adminPath := "/admin"
	if p := os.Getenv("ADMIN_PATH"); p != "" {
		p = "/" + strings.Trim(p, "/")
		adminPath = p
	}
	// 静态文件服务
	r.Static(adminPath, "./web/out")
	r.Static("/_next", "./web/out/_next")
	r.Static("/public", "./web/public")

	adminHandler := handlers.NewAdminHandler(model)
	// 管理页面路由，需要认证
	admin := r.Group("/api", handlers.BasicAuth())
	{
		// 管理页面首页
		// admin.GET("/", handlers.AdminIndex)

		// 链接配置 API
		admin.GET("/link-configs", adminHandler.GetLinkConfigs)
		admin.POST("/link-configs", adminHandler.CreateLinkConfig)
		admin.PUT("/link-configs/:id", adminHandler.UpdateLinkConfig)
		admin.DELETE("/link-configs/:id", adminHandler.DeleteLinkConfig)

		// Merge 配置 API
		admin.GET("/merge-configs", adminHandler.GetMergeConfigs)
		admin.POST("/merge-configs", adminHandler.CreateMergeConfig)
		admin.PUT("/merge-configs/:id", adminHandler.UpdateMergeConfig)
		admin.DELETE("/merge-configs/:id", adminHandler.DeleteMergeConfig)
	}
}

// func parseConfig(configFileName string) ([]*models.ConfigItem, error) {
// 	// 根据配置名称读取文件
// 	configFilePath := filepath.Join(configDir, configFileName)
// 	configContent, err := os.ReadFile(configFilePath)
// 	if err != nil {
// 		return nil, fmt.Errorf("无法读取配置文件: %v", err)
// 	}
// 	log.Infof("成功读取配置文件: %s", configFileName)
// 	var configItemArr []*models.ConfigItem
// 	err = yaml.Unmarshal(configContent, &configItemArr)
// 	if err != nil {
// 		return nil, fmt.Errorf("解析配置文件失败: %v", err)
// 	}
// 	return configItemArr, nil
// }

// func processConfig(configItemArr []*models.ConfigItem) ([]*models.ConfigItem, error) {
// 	for i, data := range configItemArr {
// 		item := data
// 		_type := item.Type
// 		if _type == "" {
// 			return nil, errors.New("type is required")
// 		}
// 		if _type == "url" {
// 			_data, err := http.Get(item.URL)
// 			if err != nil {
// 				return nil, fmt.Errorf("获取URL失败,url %s, %v", item.URL, err)
// 			}
// 			defer _data.Body.Close()
// 			dataBytes, err := io.ReadAll(_data.Body)
// 			if err != nil {
// 				return nil, fmt.Errorf("读取URL内容失败,url %s, %v", item.URL, err)
// 			}
// 			var dataMap map[string]any
// 			err = yaml.Unmarshal(dataBytes, &dataMap)
// 			if err != nil {
// 				return nil, fmt.Errorf("解析URL内容失败,url %s, %v", item.URL, err)
// 			}
// 			item.Data = dataMap
// 			configItemArr[i] = item
// 		}
// 	}
// 	return configItemArr, nil
// }

// func initLog(_filepath string) {
// 	level := logrus.InfoLevel
// 	envLevel := os.Getenv("LOG_LEVEL")
// 	v, err := logrus.ParseLevel(envLevel)
// 	if err == nil {
// 		level = v
// 	}
// 	logrus.SetLevel(level)

// 	lumberjackLogger := &lumberjack.Logger{
// 		// Log file abbsolute path, os agnostic
// 		Filename:   filepath.ToSlash(_filepath),
// 		MaxSize:    10, // MB
// 		MaxBackups: 5,
// 		MaxAge:     30, // days
// 		//Compress:   true, // disabled by default
// 	}
// 	// Fork writing into two outputs
// 	multiWriter := io.MultiWriter(os.Stdout, lumberjackLogger)
// 	logrus.SetOutput(multiWriter)
// }
