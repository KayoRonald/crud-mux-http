package models

import (
	"time"

	"gorm.io/gorm"
)

type Tasks struct {
	gorm.Model
	ID           string `json:"id" gorm:"type:uuid;primaryKey"`
	Name         string `json:"name" gorm:"not null"`
	Descripition string `json:"descripition" gorm:"not null"`
	Done         bool   `json:"done"`
	CreateAt     time.Time
	UpdateAt     time.Time
	DeleteAt     gorm.DeletedAt `gorm:"index"`
}
