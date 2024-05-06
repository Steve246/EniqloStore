package model

import "gorm.io/gorm"

type Staff struct {
	gorm.Model
	UserID      string `json:"userID"`
	Name        string `gorm:"unique" json:"userName"`
	PhoneNumber string `json:"phoneNumber"`
}

type Customer struct {
	gorm.Model
	UserID      string `json:"userID"`
	Name        string `gorm:"unique" json:"userName"`
	PhoneNumber string `json:"phoneNumber"`
}
