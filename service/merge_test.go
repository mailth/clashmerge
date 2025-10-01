package service

import (
	"clashmerge/models"
	"fmt"
	"os"
	"testing"

	myyaml "clashmerge/lib/yaml"

	"github.com/stretchr/testify/assert"
)

func TestPrependRules(t *testing.T) {
	// 读取基础配置
	yamlFile := "testdata/test3-in.yaml"
	data, err := readTestInFile(yamlFile)
	if err != nil {
		t.Fatalf("readTestInFile failed: %v", err)
	}
	// 定义要添加的规则
	rulesYAML := []string{
		"DOMAIN-SUFFIX,bilibili.com,DIRECT",
		"DOMAIN-SUFFIX,zhihu.com,DIRECT",
	}

	// 应用 prepend-rules
	actual, err := getMergeService().applyPrependRules(data, rulesYAML)
	if err != nil {
		t.Fatalf("applyPrependRules failed: %v", err)
	}
	actualYAML, err := myyaml.MarshalIndent(actual, 2)
	if err != nil {
		t.Fatalf("marshal file %s failed: %v", yamlFile, err)
	}
	os.WriteFile("testdata/tmp.yaml", actualYAML, 0644)

	// 读取预期结果
	yamlFile = "testdata/test3-out-prepend-rules.yaml"
	res, err := os.ReadFile(yamlFile)
	if err != nil {
		t.Fatalf("read file %s failed: %v", yamlFile, err)
	}

	assert.Equal(t, string(res), string(actualYAML))
}

func TestPrependProxies(t *testing.T) {
	// 读取基础配置
	yamlFile := "testdata/test3-in.yaml"
	data, err := readTestInFile(yamlFile)
	if err != nil {
		t.Fatalf("unmarshal file %s failed: %v", yamlFile, err)
	}

	// 定义要添加的代理
	proxies := []models.Proxy{
		{
			"name":     "proxy3",
			"type":     "ss",
			"server":   "server3.com",
			"port":     443,
			"cipher":   "aes-256-gcm",
			"password": "password3",
		},
	}

	// 应用 prepend-proxies
	actual, err := getMergeService().applyPrependProxies(data, proxies)
	if err != nil {
		t.Fatalf("applyPrependProxies failed: %v", err)
	}
	actualYAML, err := myyaml.MarshalIndent(actual, 2)
	if err != nil {
		t.Fatalf("marshal file %s failed: %v", yamlFile, err)
	}
	os.WriteFile("testdata/tmp.yaml", actualYAML, 0644)

	// 读取预期结果
	yamlFile = "testdata/test3-out-prepend-proxies.yaml"
	res, err := os.ReadFile(yamlFile)
	if err != nil {
		t.Fatalf("read file %s failed: %v", yamlFile, err)
	}

	assert.Equal(t, string(res), string(actualYAML))
}

func TestPrependProxyGroups(t *testing.T) {
	// 读取基础配置
	yamlFile := "testdata/test3-in.yaml"
	data, err := readTestInFile(yamlFile)
	if err != nil {
		t.Fatalf("read file %s failed: %v", yamlFile, err)
	}

	// 定义要添加的代理组
	proxyGroups := []models.ProxyGroup{
		{
			Name:    "Auto",
			Type:    "url-test",
			Proxies: []string{"direct-proxy", "Proxy"},
		},
	}

	// 应用 prepend-proxy-groups
	actual, err := getMergeService().applyPrependProxyGroups(data, proxyGroups)
	if err != nil {
		t.Fatalf("applyPrependProxyGroups failed: %v", err)
	}
	actualYAML, err := myyaml.MarshalIndent(actual, 2)
	if err != nil {
		t.Fatalf("marshal file %s failed: %v", yamlFile, err)
	}
	os.WriteFile("testdata/tmp.yaml", actualYAML, 0644)

	// 读取预期结果
	yamlFile = "testdata/test3-out-prepend-proxy-groups.yaml"
	res, err := os.ReadFile(yamlFile)
	if err != nil {
		t.Fatalf("read file %s failed: %v", yamlFile, err)
	}

	// 比较结果
	assert.Equal(t, string(res), string(actualYAML))
}

func readTestInFile(yamlFile string) (*myyaml.OrderedMap, error) {
	res, err := os.ReadFile(yamlFile)
	if err != nil {
		return nil, fmt.Errorf("read file %s failed: %v", yamlFile, err)
	}

	var data myyaml.OrderedMap
	err = myyaml.Unmarshal(res, &data)
	if err != nil {
		return nil, fmt.Errorf("unmarshal file %s failed: %v", yamlFile, err)
	}
	return &data, nil
}

func getMergeService() *MergeService {
	model, err := models.NewDB()
	if err != nil {
		panic(err)
	}
	return NewMergeService(model)
}
