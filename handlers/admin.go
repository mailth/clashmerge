package handlers

import (
	"fmt"
	"strconv"

	"clashmerge/lib/yaml"
	"clashmerge/models"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	model *models.Model
}

func NewAdminHandler(model *models.Model) *AdminHandler {
	return &AdminHandler{model: model}
}

// ===== 链接配置 API =====

// 获取所有链接配置
func (h *AdminHandler) GetLinkConfigs(c *gin.Context) {
	var configs []models.LinkConfig
	if err := h.model.DB().Find(&configs).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, configs)
}

// 创建链接配置
func (h *AdminHandler) CreateLinkConfig(c *gin.Context) {
	var config models.LinkConfig
	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.model.DB().Create(&config).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, config)
}

// 更新链接配置
func (h *AdminHandler) UpdateLinkConfig(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid ID"})
		return
	}

	var config models.LinkConfig
	if err := h.model.DB().Raw("SELECT * FROM link_configs WHERE id = ?", id).Scan(&config).Error; err != nil {
		c.JSON(404, gin.H{"error": "config not found"})
		return
	}

	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.model.DB().Save(&config).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, config)
}

// 删除链接配置
func (h *AdminHandler) DeleteLinkConfig(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid ID"})
		return
	}

	if err := h.model.DB().Delete(&models.LinkConfig{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "deleted successfully"})
}

// ===== Merge 配置 API =====

// 获取所有 Merge 配置
func (h *AdminHandler) GetMergeConfigs(c *gin.Context) {
	var configs []models.MergeConfig
	if err := h.model.DB().Find(&configs).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	reqConfigs := make([]MergeConfigReq, 0)
	for _, config := range configs {
		reqConfig, err := mergeConfigModalToReq(&config)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		reqConfigs = append(reqConfigs, *reqConfig)
	}
	c.JSON(200, reqConfigs)
}

// 创建 Merge 配置
func (h *AdminHandler) CreateMergeConfig(c *gin.Context) {
	var config MergeConfigReq
	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	mergeConfig, err := mergeConfigReqToModal(&config)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.model.DB().Create(&mergeConfig).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, config)
}

// 更新 Merge 配置
func (h *AdminHandler) UpdateMergeConfig(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid ID"})
		return
	}

	var config MergeConfigReq
	if err := h.model.DB().Raw("SELECT * FROM merge_configs WHERE id = ?", id).Scan(&config).Error; err != nil {
		c.JSON(404, gin.H{"error": "config not found"})
		return
	}

	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	mergeConfig, err := mergeConfigReqToModal(&config)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.model.DB().Save(&mergeConfig).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, config)
}

// 删除 Merge 配置
func (h *AdminHandler) DeleteMergeConfig(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid ID"})
		return
	}

	if err := h.model.DB().Delete(&models.MergeConfig{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "deleted successfully"})
}

type MergeConfigReq struct {
	ID          uint   `json:"id"`
	Name        string `json:"name" gorm:"uniqueIndex"`
	Description string `json:"description"`
	Rules       string `json:"rules" gorm:"type:text"`        // YAML 格式的规则
	Proxies     string `json:"proxies" gorm:"type:text"`      // YAML 格式的代理
	ProxyGroups string `json:"proxy_groups" gorm:"type:text"` // 代理组
}

func mergeConfigReqToModal(m *MergeConfigReq) (*models.MergeConfig, error) {
	rules := make([]string, 0)
	proxies := make([]models.Proxy, 0)
	proxyGroups := make([]models.ProxyGroup, 0)
	err := yaml.Unmarshal([]byte(m.Rules), &rules)
	if err != nil {
		return nil, fmt.Errorf("unmarshal rules failed: %v", err)
	}
	err = yaml.Unmarshal([]byte(m.Proxies), &proxies)
	if err != nil {
		return nil, fmt.Errorf("unmarshal proxies failed: %v", err)
	}
	err = yaml.Unmarshal([]byte(m.ProxyGroups), &proxyGroups)
	if err != nil {
		return nil, fmt.Errorf("unmarshal proxy groups failed: %v", err)
	}
	return &models.MergeConfig{
		BaseModel:   models.BaseModel{ID: uint(m.ID)},
		Name:        m.Name,
		Description: m.Description,
		Rules:       rules,
		Proxies:     proxies,
		ProxyGroups: proxyGroups,
	}, nil
}

func mergeConfigModalToReq(m *models.MergeConfig) (*MergeConfigReq, error) {
	rules, err := yaml.Marshal(m.Rules)
	if err != nil {
		return nil, err
	}
	proxies, err := yaml.Marshal(m.Proxies)
	if err != nil {
		return nil, err
	}
	proxyGroups, err := yaml.Marshal(m.ProxyGroups)
	if err != nil {
		return nil, err
	}
	return &MergeConfigReq{
		ID:          m.BaseModel.ID,
		Name:        m.Name,
		Description: m.Description,
		Rules:       string(rules),
		Proxies:     string(proxies),
		ProxyGroups: string(proxyGroups),
	}, nil
}
