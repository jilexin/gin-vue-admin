package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var Email = new(email)

type email struct{}

func (r *email) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := private.Group("emailplus").Use(middleware.OperationRecord())
		group.POST("emailTest", apiEmail.EmailTest) // 发送测试邮件
		group.POST("sendEmail", apiEmail.SendEmail) // 发送邮件
	}
}
