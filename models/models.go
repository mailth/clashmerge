package models

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var (
	_db *gorm.DB
)

func NewDB() (*Model, error) {
	// 设置日志输出前缀
	dataDir := "/data"
	if p := os.Getenv("DATA_DIR"); p != "" {
		dataDir = p
	}

	// 初始化数据库
	var err error
	dbPath := filepath.Join(dataDir, "clashmerge.db")
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}

	// 自动迁移数据库表
	err = db.AutoMigrate(&LinkConfig{}, &MergeConfig{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}

	_db = db
	return &Model{_db: db}, nil
}

type Model struct {
	_db *gorm.DB
}

func (m *Model) DB() *gorm.DB {
	return m._db
}

func (m *Model) GetLinkConfig(name string) (*LinkConfig, error) {
	var linkConfig LinkConfig
	if err := m.DB().Where("name = ?", name).First(&linkConfig).Error; err != nil {
		return nil, err
	}
	return &linkConfig, nil
}

func (m *Model) GetMergeConfig(id uint) (*MergeConfig, error) {
	var mergeConfig MergeConfig
	if err := m.DB().First(&mergeConfig, id).Error; err != nil {
		return nil, err
	}
	return &mergeConfig, nil
}
