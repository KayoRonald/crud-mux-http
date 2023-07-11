package models

import (
	"time"

	"gorm.io/gorm"
)

type Tasks struct {
	gorm.Model
	ID           string `gorm:"type:uuid;primaryKey"`
	Name         string `gorm:"not null"`
	Descripition string `gorm:"not null"`
	Done         bool
	CreateAt     time.Time
	UpdateAt     time.Time
	DeleteAt     gorm.DeletedAt `gorm:"index"`
}
