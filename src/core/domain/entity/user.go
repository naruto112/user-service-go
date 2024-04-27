package entity

import (
	"time"
)

type User struct {
	ID        uint
	Name      string
	Email     string
	CreatedAt time.Time
	UpdateAt  time.Time
}
