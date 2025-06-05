package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/email/api"
	"github.com/gin-gonic/gin"
)

type EmailRouter struct{}

func (s *EmailRouter) InitEmailRouter(Router *gin.RouterGroup) {
	emailRouter := Router.Use(middleware.OperationRecord())
	EmailApi := api.ApiGroupApp.EmailApi
	{
		// 邮件发送相关
		emailRouter.POST("emailTest", EmailApi.EmailTest) // 发送测试邮件
		emailRouter.POST("sendEmail", EmailApi.SendEmail) // 发送邮件

		// 草稿管理相关
		emailRouter.POST("saveDraft", EmailApi.SaveEmailDraft)           // 保存草稿
		emailRouter.PUT("updateDraft/:id", EmailApi.UpdateEmailDraft)    // 更新草稿
		emailRouter.GET("getDrafts", EmailApi.GetEmailDrafts)            // 获取草稿列表
		emailRouter.GET("getDraft/:id", EmailApi.GetEmailDraftDetail)    // 获取草稿详情
		emailRouter.DELETE("deleteDraft/:id", EmailApi.DeleteEmailDraft) // 删除草稿
	}
}
