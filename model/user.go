package model

import "gorm.io/gorm"

type Staff struct {
	gorm.Model
	UserID      string `json:"userID"`
	Name        string `gorm:"unique" json:"userName"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

type StaffResponse struct {
	UserID      string `json:"userID"`
	Name        string `gorm:"unique" json:"userName"`
	PhoneNumber string `json:"phoneNumber"`
	AccessToken string `json:"accesToken"`
}

type Customer struct {
	gorm.Model
	UserID      string `json:"userID"`
	Name        string `gorm:"unique" json:"userName"`
	PhoneNumber string `json:"phoneNumber"`
}
