package models

import (
	"time"
)

type User struct {
	Id        uint      `gorm:"primary_key" json:"id"`
	Username  string    `gorm:"type:varchar(300)" json:"username" binding:"required"`
	Password  string    `gorm:"type:varchar(300)" json:"password" binding:"required"`
	Role      string    `gorm:"type:varchar(300)" json:"role" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
