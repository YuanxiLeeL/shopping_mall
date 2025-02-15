package models

import "gorm.io/gorm"

type ReplyComment struct {
	gorm.Model
	UserID  uint   `json:"user_id" binding:"required" gorm:"index"`
	GoodID  uint   `json:"good_id" binding:"required" gorm:"index"`
	Content string `json:"content" binding:"required"`
	ReplyTo uint   `json:"reply_to" gorm:"default:0"` // 回复的评论ID，0表示不是回复
}
