package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         uint           `json:"id" gorm:"primary_key"`
	Email      string         `json:"email"`
	First_name string         `json:"first_name"`
	Last_name  string         `json:"last_name"`
	Avatar     string         `json:"avatar"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

func (u *User) Delete(db *gorm.DB) error {
	return db.Model(u).Update("deleted_at", time.Now()).Error
}
