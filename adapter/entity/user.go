package entity

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey,autoIncrement"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"not null;unique`
	CreatedAt time.Time `gorm: autoCreateTime`
	UpdateAt  time.Time `gorm: autoUpdateTime`
}
