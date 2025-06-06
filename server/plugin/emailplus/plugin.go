package emailplus

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/emailplus/initialize"
	interfaces "github.com/flipped-aurora/gin-vue-admin/server/utils/plugin/v2"
	"github.com/gin-gonic/gin"
)

// 确保实现了v2插件接口
var _ interfaces.Plugin = (*plugin)(nil)

// Plugin 导出的插件实例
var Plugin = new(plugin)

type plugin struct{}

// Register 实现v2插件接口，注册插件到gin引擎
func (p *plugin) Register(group *gin.Engine) {
	ctx := context.Background()

	// 配置初始化 - 如果需要配置文件支持
	// initialize.Viper()

	// API数据初始化 - 自动注册API
	initialize.Api(ctx)

	// 菜单数据初始化 - 自动注册菜单
	initialize.Menu(ctx)

	// 数据库初始化 - 自动迁移表结构
	initialize.Gorm(ctx)

	// 路由初始化 - 注册API路由
	initialize.Router(group)
}
