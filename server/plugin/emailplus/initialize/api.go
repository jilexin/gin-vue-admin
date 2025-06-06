package initialize

import (
	"context"

	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

// Api 初始化email插件API
func Api(ctx context.Context) {
	entities := []model.SysApi{
		{
			Path:        "/emailplus/emailTest",
			Description: "发送测试邮件",
			ApiGroup:    "邮件管理",
			Method:      "POST",
		},
		{
			Path:        "/emailplus/sendEmail",
			Description: "发送邮件",
			ApiGroup:    "邮件管理",
			Method:      "POST",
		},
	}
	utils.RegisterApis(entities...)
}
