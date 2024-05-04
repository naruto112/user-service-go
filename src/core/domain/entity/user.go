package entity

import (
	"time"
)

type User struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdateAt  time.Time
}
