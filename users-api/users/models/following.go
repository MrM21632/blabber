package models

import "time"

type Following struct {
	FollowerID string    `json:"follower_id" gorm:"primaryKey;autoincrement:false"`
	FollowedID string    `json:"followed_id" gorm:"primaryKey;autoincrement:false"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time

	Follower User `gorm:"foreignKey:FollowerID;constraint:OnDelete:CASCADE"`
	Followed User `gorm:"foreignKey:FollowedID;constraint:OnDelete:CASCADE"`
}
