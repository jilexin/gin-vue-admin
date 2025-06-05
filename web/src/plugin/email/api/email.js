import service from '@/utils/request'
// @Tags System
// @Summary 发送测试邮件
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /email/emailTest [post]
export const emailTest = (data) => {
  return service({
    url: '/email/emailTest',
    method: 'post',
    data
  })
}

// @Tags System
// @Summary 发送邮件
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body email_response.Email true "发送邮件必须的参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /email/sendEmail [post]
export const sendEmail = (data) => {
  return service({
    url: '/email/sendEmail',
    method: 'post',
    data
  })
}

// @Tags System
// @Summary 保存邮件草稿
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body EmailDraftRequest true "保存草稿必须的参数"
// @Success 200 {object} model.EmailDraft "{"success":true,"data":{},"msg":"保存成功"}"
// @Router /email/saveDraft [post]
export const saveEmailDraft = (data) => {
  return service({
    url: '/email/saveDraft',
    method: 'post',
    data
  })
}

// @Tags System
// @Summary 更新邮件草稿
// @Security ApiKeyAuth
// @Produce  application/json
// @Param id path int true "草稿ID"
// @Param data body EmailDraftRequest true "更新草稿必须的参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /email/updateDraft/{id} [put]
export const updateEmailDraft = (id, data) => {
  return service({
    url: `/email/updateDraft/${id}`,
    method: 'put',
    data
  })
}

// @Tags System
// @Summary 获取邮件草稿列表
// @Security ApiKeyAuth
// @Produce  application/json
// @Param page query int false "页码"
// @Param pageSize query int false "每页数量"
// @Success 200 {object} response.PageResult "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /email/getDrafts [get]
export const getEmailDrafts = (params) => {
  return service({
    url: '/email/getDrafts',
    method: 'get',
    params
  })
}

// @Tags System
// @Summary 获取邮件草稿详情
// @Security ApiKeyAuth
// @Produce  application/json
// @Param id path int true "草稿ID"
// @Success 200 {object} model.EmailDraft "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /email/getDraft/{id} [get]
export const getEmailDraftDetail = (id) => {
  return service({
    url: `/email/getDraft/${id}`,
    method: 'get'
  })
}

// @Tags System
// @Summary 删除邮件草稿
// @Security ApiKeyAuth
// @Produce  application/json
// @Param id path int true "草稿ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /email/deleteDraft/{id} [delete]
export const deleteEmailDraft = (id) => {
  return service({
    url: `/email/deleteDraft/${id}`,
    method: 'delete'
  })
}
