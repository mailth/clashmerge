package handlers

import (
	myyaml "clashmerge/lib/yaml"
	"clashmerge/models"
	"clashmerge/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	keepHeaderKeys = []string{"Subscription-Userinfo", "Content-Disposition"}
)

type MergeHandler struct {
	model        *models.Model
	mergeService *service.MergeService
}

func NewMergeHandler(model *models.Model, mergeService *service.MergeService) *MergeHandler {
	return &MergeHandler{model: model, mergeService: mergeService}
}

func (h *MergeHandler) HandleMerge(c *gin.Context) {
	linkName := c.Query("name")
	if linkName == "" {
		c.JSON(400, gin.H{"error": "name is required"})
		return
	}
	logrus.Infof("handle merge, linkName: %s", linkName)

	res, headers, err := h.mergeService.ProcessConfig(linkName)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	en, err := myyaml.MarshalIndent(res, 2)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// 设置文件下载头
	// filename := linkName + ".yaml"
	// c.Header("Content-Disposition", "attachment; filename=\""+filename+"\"")
	c.Header("Content-Type", "application/x-yaml")
	for _, v := range keepHeaderKeys {
		if vv, ok := headers[v]; ok {
			c.Header(v, vv[0])
		}
	}
	c.String(200, string(en))
}
