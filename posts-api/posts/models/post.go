package models

import "time"

type Post struct {
	ID        string    `json:"id" gorm:"primaryKey;autoincrement:false"`
	ParentID  *string   `json:"parent_post_id"`
	UserID    string    `json:"user_id" gorm:"unique;not null"`
	Contents  string    `json:"contents" gorm:"size:200;not null"`
	Likes     uint32    `json:"likes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Replies []Post `gorm:"foreignkey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Post) TableName() string {
	return "post"
}
