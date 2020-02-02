package models

import "github.com/jinzhu/gorm"

type Vote struct {
	gorm.Model
	CommentID uint `json: "commentId" gorm:"not null"`
	UserID uint `json:"userId" gorm:"not null"`
	value bool `json:"value" gorm:"not null"`

}