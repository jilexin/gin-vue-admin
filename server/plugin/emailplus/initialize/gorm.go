package initialize

import (
	"context"
)

// Gorm 初始化email插件数据库表
func Gorm(ctx context.Context) {
	// email插件目前没有需要自动迁移的数据表
	// 如果需要添加表模型，请在这里进行AutoMigrate
}
