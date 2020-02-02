package models

import "github.com/jinzhu/gorm"

type Usuario struct {
	gorm.Model
	NickName   string `json:"nickname" gorm:"not null"`
	FirstName  string `json:"firstname" gorm:"not null"`
	SecondName string `json:"secondname" gorm:"not null"`
	LastName   string `json:"lastname" gorm:"not null"`
	PIN        uint16 `json:"pin,omitempty" gorm:"not null"`
	Picture    string `json:"picture,omitempty"`
}
