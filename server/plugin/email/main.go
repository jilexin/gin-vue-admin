package email

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/email/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/email/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/email/router"
	"github.com/gin-gonic/gin"
)

type emailPlugin struct{}

func CreateEmailPlug(To, From, Host, Secret, Nickname string, Port int, IsSSL bool, IsLoginAuth bool) *emailPlugin {
	global.GlobalConfig.To = To
	global.GlobalConfig.From = From
	global.GlobalConfig.Host = Host
	global.GlobalConfig.Secret = Secret
	global.GlobalConfig.Nickname = Nickname
	global.GlobalConfig.Port = Port
	global.GlobalConfig.IsSSL = IsSSL
	global.GlobalConfig.IsLoginAuth = IsLoginAuth
	return &emailPlugin{}
}

func (*emailPlugin) Register(group *gin.RouterGroup) {
	ctx := context.Background()
	router.RouterGroupApp.InitEmailRouter(group)
	initialize.Gorm(ctx)
}

func (*emailPlugin) RouterPath() string {
	return "email"
}
