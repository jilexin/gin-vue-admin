package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/email/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/email/utils"
)

type EmailService struct{}

//@author: [maplepie](https://github.com/maplepie)
//@function: EmailTest
//@description: 发送邮件测试
//@return: err error

func (e *EmailService) EmailTest() (err error) {
	subject := "test"
	body := "test"
	err = utils.EmailTest(subject, body)
	return err
}

//@author: [maplepie](https://github.com/maplepie)
//@function: EmailTest
//@description: 发送邮件测试
//@return: err error
//@params to string 	 收件人
//@params subject string   标题（主题）
//@params body  string 	 邮件内容

func (e *EmailService) SendEmail(to, subject, body string) (err error) {
	err = utils.Email(to, subject, body)
	return err
}

//@author: [您的名称]
//@function: SaveEmailDraft
//@description: 保存邮件草稿
//@return: draft model.EmailDraft, err error
//@params to string 	 收件人
//@params subject string   标题（主题）
//@params body  string 	 邮件内容
//@params title string 	 草稿标题
//@params userID uint   用户ID

func (e *EmailService) SaveEmailDraft(to, subject, body, title string, userID uint) (draft model.EmailDraft, err error) {
	// 如果没有提供草稿标题，使用邮件主题作为标题
	if title == "" {
		title = subject
		if title == "" {
			title = "未命名草稿"
		}
	}

	draft = model.EmailDraft{
		To:      to,
		Subject: subject,
		Body:    body,
		Title:   title,
		UserID:  userID,
	}

	err = global.GVA_DB.Create(&draft).Error
	return draft, err
}

//@author: [您的名称]
//@function: UpdateEmailDraft
//@description: 更新邮件草稿
//@return: err error
//@params id uint   草稿ID
//@params to string 	 收件人
//@params subject string   标题（主题）
//@params body  string 	 邮件内容
//@params title string 	 草稿标题

func (e *EmailService) UpdateEmailDraft(id uint, to, subject, body, title string) (err error) {
	// 如果没有提供草稿标题，使用邮件主题作为标题
	if title == "" {
		title = subject
		if title == "" {
			title = "未命名草稿"
		}
	}

	updates := map[string]interface{}{
		"to":      to,
		"subject": subject,
		"body":    body,
		"title":   title,
	}

	err = global.GVA_DB.Model(&model.EmailDraft{}).Where("id = ?", id).Updates(updates).Error
	return err
}

//@author: [您的名称]
//@function: GetEmailDrafts
//@description: 获取用户的邮件草稿列表
//@return: drafts []model.EmailDraft, total int64, err error
//@params userID uint   用户ID
//@params page int   页码
//@params pageSize int   每页数量

func (e *EmailService) GetEmailDrafts(userID uint, page, pageSize int) (drafts []model.EmailDraft, total int64, err error) {
	limit := pageSize
	offset := pageSize * (page - 1)

	db := global.GVA_DB.Model(&model.EmailDraft{}).Where("user_id = ?", userID)

	// 获取总数
	err = db.Count(&total).Error
	if err != nil {
		return drafts, total, err
	}

	// 获取草稿列表，按更新时间倒序
	err = db.Order("updated_at desc").Limit(limit).Offset(offset).Find(&drafts).Error

	return drafts, total, err
}

//@author: [您的名称]
//@function: GetEmailDraftByID
//@description: 根据ID获取邮件草稿详情
//@return: draft model.EmailDraft, err error
//@params id uint   草稿ID
//@params userID uint   用户ID（用于权限验证）

func (e *EmailService) GetEmailDraftByID(id, userID uint) (draft model.EmailDraft, err error) {
	err = global.GVA_DB.Where("id = ? AND user_id = ?", id, userID).First(&draft).Error
	return draft, err
}

//@author: [您的名称]
//@function: DeleteEmailDraft
//@description: 删除邮件草稿
//@return: err error
//@params id uint   草稿ID
//@params userID uint   用户ID（用于权限验证）

func (e *EmailService) DeleteEmailDraft(id, userID uint) (err error) {
	err = global.GVA_DB.Where("id = ? AND user_id = ?", id, userID).Delete(&model.EmailDraft{}).Error
	return err
}
