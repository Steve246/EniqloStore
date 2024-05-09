package model

import "time"

type ProductList struct {
	ID          int       `gorm:"primaryKey"`
	Name        string    `gorm:"size:30;not null"`
	SKU         string    `gorm:"size:30;not null"`
	Category    string    `gorm:"size:60;not null"`
	ImageURL    string    `gorm:"size:100;not null"`
	Notes       string    `gorm:"size:200;not null"`
	Price       int       `gorm:"not null;check:price >= 0"`
	Stock       int       `gorm:"not null;check:stock >= 0 AND stock <= 100000"`
	Location    string    `gorm:"size:200;not null"`
	IsAvailable bool      `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	DeletedAt   time.Time `gorm:"index"`
}
