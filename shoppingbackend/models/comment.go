package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserName  string   `json:"username" gorm:"index"`
	GoodName  string   `json:"goodname" gorm:"index"`
	Content string `json:"content" binding:"required"`
	// ReplyTo uint   `json:"reply_to" gorm:"default:0"` // 回复的评论ID，0表示不是回复
	// Rating  uint8  `json:"rating" binding:"omitempty,numeric,min=0,max=5"` // 评分，范围从0到5，可选
}




