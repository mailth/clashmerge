package service

import (
	"fmt"
	"io"
	"net/http"

	"clashmerge/models"

	myyaml "clashmerge/lib/yaml"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type MergeService struct {
	model *models.Model
}

func NewMergeService(model *models.Model) *MergeService {
	return &MergeService{model: model}
}

// HandleMerge 处理配置合并请求

// // 新的合并函数，支持通过链接配置和 Merge 配置进行合并
// func doMergeWithConfigs(linkConfigName string) (map[string]any, error) {
// 	// 查找链接配置
// 	var linkConfig models.LinkConfig
// 	if err := DB.Where("name = ?", linkConfigName).First(&linkConfig).Error; err != nil {
// 		return nil, fmt.Errorf("链接配置不存在: %v", err)
// 	}

// 	// 获取原始 Clash 配置
// 	data, err := fetchClashConfig(linkConfig.ClashURL)
// 	if err != nil {
// 		return nil, fmt.Errorf("获取 Clash 配置失败: %v", err)
// 	}

// 	// 如果有关联的 Merge 配置，则应用它
// 	if linkConfig.MergeConfigID > 0 {
// 		var mergeConfig models.MergeConfig
// 		if err := DB.First(&mergeConfig, linkConfig.MergeConfigID).Error; err != nil {
// 			return nil, fmt.Errorf("Merge 配置不存在: %v", err)
// 		}

// 		data, err = applyMergeConfig(data, &mergeConfig)
// 		if err != nil {
// 			return nil, fmt.Errorf("应用 Merge 配置失败: %v", err)
// 		}
// 	}

// 	return data, nil
// }

// // 获取配置目录
// func getConfigDir() string {
// 	configDir := os.Getenv("CONFIG_DIR")
// 	if configDir == "" {
// 		configDir = "/data"
// 	}
// 	return configDir
// }

// func parseConfig(configFileName string) ([]*models.ConfigItem, error) {
// 	// 根据配置名称读取文件
// 	configDir := getConfigDir()
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

func (s *MergeService) ProcessConfig(name string) (*myyaml.OrderedMap, http.Header, error) {
	// 1. get link config
	linkConfig, err := s.model.GetLinkConfig(name)
	if err != nil {
		return nil, nil, err
	}
	// 2. get link config detail
	_data, err := http.Get(linkConfig.ClashURL)
	if err != nil {
		return nil, nil, fmt.Errorf("get URL failed,url %s, %v", linkConfig.ClashURL, err)
	}
	defer _data.Body.Close()
	dataBytes, err := io.ReadAll(_data.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("read URL content failed,url %s, %v", linkConfig.ClashURL, err)
	}
	dataMap := &myyaml.OrderedMap{}
	err = yaml.Unmarshal(dataBytes, dataMap)
	if err != nil {
		return nil, nil, fmt.Errorf("parse URL content failed,url %s, %v", linkConfig.ClashURL, err)
	}

	logrus.Infof("heand:%v", _data.Header)
	// without merge config
	if linkConfig.MergeConfigID <= 0 {
		logrus.Warnf("no merge config, mergeConfigID: %d", linkConfig.MergeConfigID)
		return dataMap, _data.Header, nil
	}

	// 3. get merge config
	mergeConfig, err := s.model.GetMergeConfig(linkConfig.MergeConfigID)
	if err != nil {
		return nil, nil, err
	}
	dataMap, err = s.applyMergeConfig(dataMap, mergeConfig)
	if err != nil {
		return nil, nil, err
	}
	return dataMap, _data.Header, nil
}

// func doMerge(data map[string]any, items []*models.ConfigItem) (map[string]any, error) {
// 	if data == nil {
// 		data = make(map[string]any)
// 	}
// 	for _, item := range items {
// 		var err error
// 		data, err = merge(data, item.Data, item.Operation)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}
// 	return data, nil
// }

// func merge(conf1, conf2 map[string]any, operation string) (map[string]any, error) {
// 	for k, v := range conf2 {
// 		_, ok := conf1[k]
// 		if !ok || operation == "replace" {
// 			conf1[k] = v
// 			continue
// 		}

// 		if operation == "merge" {
// 			_val, err := mergeVal(conf1[k], v)
// 			if err != nil {
// 				return nil, err
// 			}
// 			conf1[k] = _val
// 		}
// 	}
// 	return conf1, nil
// }

// func mergeVal(val1, val2 any) (any, error) {
// 	if val2 == nil {
// 		return val1, nil
// 	}

// 	switch v1 := val1.(type) {
// 	case []any:
// 		switch v2 := val2.(type) {
// 		case []any:
// 			res := append(v1, v2...)
// 			res = dedup(res)
// 			return res, nil
// 		default:
// 			return nil, errors.New("append operation not supported for non-list types")
// 		}
// 	case map[string]any:
// 		switch v2 := val2.(type) {
// 		case map[string]any:
// 			for k, v := range v2 {
// 				v1[k] = v
// 			}
// 			return v1, nil
// 		default:
// 			return nil, errors.New("append operation not supported for non-map types")
// 		}
// 	default:
// 		return val2, nil
// 	}
// }

// func dedup[T comparable](list []T) []T {
// 	seen := make(map[T]struct{})
// 	result := make([]T, 0)

// 	for _, item := range list {
// 		if _, exists := seen[item]; !exists {
// 			seen[item] = struct{}{}
// 			result = append(result, item)
// 		}
// 	}

// 	return result
// }

// 获取 Clash 配置
// func (s *MergeService) fetchClashConfig(url string) (*myyaml.OrderedMap, error) {
// 	// 这里复用原有的 HTTP 获取逻辑
// 	configItem := &models.ConfigItem{
// 		Type: "url",
// 		URL:  url,
// 	}

// 	items := []*models.ConfigItem{configItem}
// 	processedItems, err := processConfig(items)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if len(processedItems) > 0 && processedItems[0].Data != nil {
// 		return processedItems[0].Data, nil
// 	}

// 	return make(map[string]yaml.Node), nil
// }

// 应用 Merge 配置
func (s *MergeService) applyMergeConfig(data *myyaml.OrderedMap, mergeConfig *models.MergeConfig) (*myyaml.OrderedMap, error) {
	var err error
	if len(mergeConfig.Rules) > 0 {
		data, err = s.applyPrependRules(data, mergeConfig.Rules)
		if err != nil {
			return nil, err
		}
	}
	if len(mergeConfig.Proxies) > 0 {
		data, err = s.applyPrependProxies(data, mergeConfig.Proxies)
		if err != nil {
			return nil, err
		}
	}
	if len(mergeConfig.ProxyGroups) > 0 {
		data, err = s.applyPrependProxyGroups(data, mergeConfig.ProxyGroups)
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}

// 应用 prepend-rules
func (s *MergeService) applyPrependRules(data *myyaml.OrderedMap, rules []string) (*myyaml.OrderedMap, error) {
	if len(rules) == 0 {
		return data, nil
	}

	// 获取现有规则
	var existingRules []string
	node, ok := data.Get("rules")
	if ok {
		err := node.Decode(&existingRules)
		if err != nil {
			return nil, fmt.Errorf("decode rules failed: %v", err)
		}
	} else {
		existingRules = []string{}
	}

	// 将新规则添加到前面
	allRules := append(rules, existingRules...)
	n := yaml.Node{}
	err := n.Encode(allRules)
	if err != nil {
		return nil, fmt.Errorf("encode rules failed: %v", err)
	}
	data.Set("rules", &n)

	return data, nil
}

// 应用 prepend-proxies
func (s *MergeService) applyPrependProxies(data *myyaml.OrderedMap, proxies []models.Proxy) (*myyaml.OrderedMap, error) {
	if len(proxies) == 0 {
		return data, nil
	}

	// 获取现有代理
	var existingProxies []myyaml.OrderedMap
	oldProxies, ok := data.Get("proxies")
	if ok {
		err := oldProxies.Decode(&existingProxies)
		if err != nil {
			return nil, fmt.Errorf("decode proxies failed: %v", err)
		}
	} else {
		existingProxies = make([]myyaml.OrderedMap, 0)
	}

	newProxies := make([]any, 0)
	for _, proxy := range proxies {
		newProxies = append(newProxies, proxy)
	}
	for _, proxy := range existingProxies {
		newProxies = append(newProxies, proxy)
	}

	n := yaml.Node{}
	err := n.Encode(newProxies)
	if err != nil {
		return nil, fmt.Errorf("encode proxies failed: %v", err)
	}
	data.Set("proxies", &n)
	return data, nil
}

func (s *MergeService) applyPrependProxyGroups(data *myyaml.OrderedMap, proxyGroups []models.ProxyGroup) (*myyaml.OrderedMap, error) {
	if len(proxyGroups) == 0 {
		return data, nil
	}

	// 获取现有代理组
	var existingProxyGroups []models.ProxyGroup
	node, ok := data.Get("proxy-groups")
	if ok {
		err := node.Decode(&existingProxyGroups)
		if err != nil {
			logrus.Error("proxy-groups 字段不是预期的类型，尝试转换为 []models.ProxyGroup")
			return nil, err
		}
	}

	newProxyGroups := append(proxyGroups, existingProxyGroups...)
	n := yaml.Node{}
	err := n.Encode(newProxyGroups)
	if err != nil {
		return nil, fmt.Errorf("encode proxy-groups failed: %v", err)
	}
	data.Set("proxy-groups", &n)
	return data, nil
}
