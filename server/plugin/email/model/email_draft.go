package model

import (
	"time"

	"gorm.io/gorm"
)

// EmailDraft 邮件草稿模型
type EmailDraft struct {
	ID        uint           `gorm:"primarykey" json:"id"`                           // 主键ID
	CreatedAt time.Time      `json:"created_at"`                                     // 创建时间
	UpdatedAt time.Time      `json:"updated_at"`                                     // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                                 // 删除时间
	To        string         `json:"to" gorm:"column:to;comment:收件人邮箱"`              // 收件人邮箱
	Subject   string         `json:"subject" gorm:"column:subject;comment:邮件主题"`     // 邮件主题
	Body      string         `json:"body" gorm:"type:text;column:body;comment:邮件内容"` // 邮件内容
	UserID    uint           `json:"user_id" gorm:"column:user_id;comment:创建用户ID"`   // 创建用户ID
	Title     string         `json:"title" gorm:"column:title;comment:草稿标题"`         // 草稿标题（用于标识）
}

// TableName 设置表名
func (EmailDraft) TableName() string {
	return "email_drafts"
}
