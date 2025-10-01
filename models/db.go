package models

import (
	"time"

	"gopkg.in/yaml.v3"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// 原有的配置项结构
type ConfigItem struct {
	Operation string               `yaml:"operation"`
	Type      string               `yaml:"type"`
	Data      map[string]yaml.Node `yaml:"data"`
	URL       string               `yaml:"url"`
}

// 链接配置模型
type LinkConfig struct {
	BaseModel
	Name          string `json:"name" gorm:"uniqueIndex"`
	ClashURL      string `json:"clash_url"`
	Description   string `json:"description"`
	MergeConfigID uint   `json:"merge_config_id"`
}

// Merge 配置模型
type MergeConfig struct {
	BaseModel
	Name        string                          `json:"name" gorm:"uniqueIndex"`
	Description string                          `json:"description"`
	Rules       datatypes.JSONSlice[string]     `json:"rules" gorm:"type:text"`        // YAML 格式的规则
	Proxies     datatypes.JSONSlice[Proxy]      `json:"proxies" gorm:"type:text"`      // YAML 格式的代理
	ProxyGroups datatypes.JSONSlice[ProxyGroup] `json:"proxy_groups" gorm:"type:text"` // 代理组
}

type ProxyGroup struct {
	Name    string   `yaml:"name"`
	Type    string   `yaml:"type"`
	Proxies []string `yaml:"proxies"`
}

type Proxy map[string]any

type BaseModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
