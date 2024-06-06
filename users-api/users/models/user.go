package models

import "time"

type User struct {
	ID        string    `json:"id" gorm:"primaryKey;autoincrement:false"`
	Handle    string    `json:"handle" gorm:"unique;not null"`
	Username  string    `json:"username"`
	Email     string    `json:"email_address" gorm:"unique;not null"`
	Password  string    `json:"password_hash"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "user"
}
