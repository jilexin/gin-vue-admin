package api

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	email_response "github.com/flipped-aurora/gin-vue-admin/server/plugin/email/model/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/email/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EmailApi struct{}

// EmailTest
// @Tags      System
// @Summary   发送测试邮件
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200  {string}  string  "{"success":true,"data":{},"msg":"发送成功"}"
// @Router    /email/emailTest [post]
func (s *EmailApi) EmailTest(c *gin.Context) {
	err := service.ServiceGroupApp.EmailTest()
	if err != nil {
		global.GVA_LOG.Error("发送失败!", zap.Error(err))
		response.FailWithMessage("发送失败", c)
		return
	}
	response.OkWithMessage("发送成功", c)
}

// SendEmail
// @Tags      System
// @Summary   发送邮件
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body      email_response.Email  true  "发送邮件必须的参数"
// @Success   200   {string}  string                "{"success":true,"data":{},"msg":"发送成功"}"
// @Router    /email/sendEmail [post]
func (s *EmailApi) SendEmail(c *gin.Context) {
	var email email_response.Email
	err := c.ShouldBindJSON(&email)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.ServiceGroupApp.SendEmail(email.To, email.Subject, email.Body)
	if err != nil {
		global.GVA_LOG.Error("发送失败!", zap.Error(err))
		response.FailWithMessage("发送失败", c)
		return
	}
	response.OkWithMessage("发送成功", c)
}

// SaveEmailDraft
// @Tags      System
// @Summary   保存邮件草稿
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body      EmailDraftRequest  true  "保存草稿必须的参数"
// @Success   200   {object}  model.EmailDraft    "{"success":true,"data":{},"msg":"保存成功"}"
// @Router    /email/saveDraft [post]
func (s *EmailApi) SaveEmailDraft(c *gin.Context) {
	type EmailDraftRequest struct {
		To      string `json:"to"`
		Subject string `json:"subject"`
		Body    string `json:"body"`
		Title   string `json:"title"`
	}

	var req EmailDraftRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 获取当前用户ID
	userID := utils.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("获取用户信息失败", c)
		return
	}

	draft, err := service.ServiceGroupApp.SaveEmailDraft(req.To, req.Subject, req.Body, req.Title, userID)
	if err != nil {
		global.GVA_LOG.Error("保存草稿失败!", zap.Error(err))
		response.FailWithMessage("保存草稿失败", c)
		return
	}

	response.OkWithDetailed(draft, "保存成功", c)
}

// UpdateEmailDraft
// @Tags      System
// @Summary   更新邮件草稿
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     id    path      int                true  "草稿ID"
// @Param     data  body      EmailDraftRequest  true  "更新草稿必须的参数"
// @Success   200   {string}  string             "{"success":true,"data":{},"msg":"更新成功"}"
// @Router    /email/updateDraft/{id} [put]
func (s *EmailApi) UpdateEmailDraft(c *gin.Context) {
	type EmailDraftRequest struct {
		To      string `json:"to"`
		Subject string `json:"subject"`
		Body    string `json:"body"`
		Title   string `json:"title"`
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("无效的ID参数", c)
		return
	}

	var req EmailDraftRequest
	err = c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = service.ServiceGroupApp.UpdateEmailDraft(uint(id), req.To, req.Subject, req.Body, req.Title)
	if err != nil {
		global.GVA_LOG.Error("更新草稿失败!", zap.Error(err))
		response.FailWithMessage("更新草稿失败", c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

// GetEmailDrafts
// @Tags      System
// @Summary   获取邮件草稿列表
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     page      query     int     false  "页码"
// @Param     pageSize  query     int     false  "每页数量"
// @Success   200       {object}  response.PageResult  "{"success":true,"data":{},"msg":"获取成功"}"
// @Router    /email/getDrafts [get]
func (s *EmailApi) GetEmailDrafts(c *gin.Context) {
	// 获取查询参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	// 参数验证
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	// 获取当前用户ID
	userID := utils.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("获取用户信息失败", c)
		return
	}

	drafts, total, err := service.ServiceGroupApp.GetEmailDrafts(userID, page, pageSize)
	if err != nil {
		global.GVA_LOG.Error("获取草稿列表失败!", zap.Error(err))
		response.FailWithMessage("获取草稿列表失败", c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     drafts,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, "获取成功", c)
}

// GetEmailDraftDetail
// @Tags      System
// @Summary   获取邮件草稿详情
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     id   path      int     true  "草稿ID"
// @Success   200  {object}  model.EmailDraft  "{"success":true,"data":{},"msg":"获取成功"}"
// @Router    /email/getDraft/{id} [get]
func (s *EmailApi) GetEmailDraftDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("无效的ID参数", c)
		return
	}

	// 获取当前用户ID
	userID := utils.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("获取用户信息失败", c)
		return
	}

	draft, err := service.ServiceGroupApp.GetEmailDraftByID(uint(id), userID)
	if err != nil {
		global.GVA_LOG.Error("获取草稿详情失败!", zap.Error(err))
		response.FailWithMessage("获取草稿详情失败", c)
		return
	}

	response.OkWithDetailed(draft, "获取成功", c)
}

// DeleteEmailDraft
// @Tags      System
// @Summary   删除邮件草稿
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     id   path      int     true  "草稿ID"
// @Success   200  {string}  string  "{"success":true,"data":{},"msg":"删除成功"}"
// @Router    /email/deleteDraft/{id} [delete]
func (s *EmailApi) DeleteEmailDraft(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("无效的ID参数", c)
		return
	}

	// 获取当前用户ID
	userID := utils.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("获取用户信息失败", c)
		return
	}

	err = service.ServiceGroupApp.DeleteEmailDraft(uint(id), userID)
	if err != nil {
		global.GVA_LOG.Error("删除草稿失败!", zap.Error(err))
		response.FailWithMessage("删除草稿失败", c)
		return
	}

	response.OkWithMessage("删除成功", c)
}
