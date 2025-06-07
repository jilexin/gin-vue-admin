package initialize

import (
	"context"

	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

// Menu 初始化email插件菜单
func Menu(ctx context.Context) {
	entities := []model.SysBaseMenu{
		{
			ParentId:  9, // 插件管理父菜单ID
			Path:      "emailplus",
			Name:      "emailplus",
			Hidden:    false,
			Component: "plugin/emailplus/view/index.vue",
			Sort:      7,
			Meta:      model.Meta{Title: "邮件管理", Icon: "message"},
		},
	}
	utils.RegisterMenus(entities...)
}
