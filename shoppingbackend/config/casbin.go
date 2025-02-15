package config

import (
	"Democratic_shopping_mall/global"
	"log"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

func InitCasbin() {
	// 初始化数据库适配器
	adapter, err := gormadapter.NewAdapter("sqlite3", "test.db", true)
	if err != nil {
		log.Fatalf("Failed to create adapter: %v", err)
	}

	// 加载模型和策略
	global.Enforcer, err = casbin.NewEnforcer("model.conf", adapter)
	if err != nil {
		log.Fatalf("Failed to create enforcer: %v", err)
	}
}
