package model

import (
	"gorm.io/gorm"
)

// type Staff struct {
// 	gorm.Model
// 	UserID      string `json:"userID"`
// 	Name        string `gorm:"unique" json:"userName"`
// 	PhoneNumber string `json:"phoneNumber"`
// 	Password    string `json:"password"`
// }

type Staff struct {
	// ID           int        `json:"id"`
	gorm.Model
	UserUniqueID string `json:"user_unique_id"`
	Name         string `json:"name"`
	Password     string `json:"password"`
	PhoneNumber  string `json:"phone_number"`
	// CreatedAt    time.Time  `json:"created_at"`
	// UpdatedAt    time.Time  `json:"updated_at"`
	// DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}

type StaffResponse struct {
	UserUniqueID string `json:"user_unique_id"`
	Name         string `gorm:"unique" json:"userName"`
	PhoneNumber  string `json:"phoneNumber"`
	AccessToken  string `json:"accesToken"`
}

type Customer struct {
	gorm.Model
	UserUniqueID string `json:"user_unique_id"`
	Name         string `gorm:"unique" json:"userName"`
	PhoneNumber  string `json:"phoneNumber"`
}
