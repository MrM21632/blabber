package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `db:"id" json:"user_id"`
	Username     string    `db:"username" json:"username"`
	Handle       string    `db:"user_handle" json:"user_handle"`
	Bio          string    `db:"user_bio" json:"user_bio"`
	Email        string    `db:"email" json:"email_address"`
	PasswordHash string    `db:"password_hash" json:"password_hash"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
	Followers    uint32    `db:"followers" json:"followers"`
	Follows      uint32    `db:"follows" json:"follows"`
}

func (User) TableName() string {
	return "user"
}

type UserFollow struct {
	FollowerID uuid.UUID `db:"follower_id" json:"follower_id"`
	FollowedID uuid.UUID `db:"followed_id" json:"followed_id"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
}

func (UserFollow) TableName() string {
	return "user_follow"
}

type UserBlock struct {
	BlockerID uuid.UUID `db:"blocker_id" json:"blocker_id"`
	BlockedID uuid.UUID `db:"blocked_id" json:"blocked_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

func (UserBlock) TableName() string {
	return "user_block"
}

type UserMute struct {
	MuterID   uuid.UUID `db:"muter_id" json:"muter_id"`
	MutedID   uuid.UUID `db:"muted_id" json:"muted_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

func (UserMute) TableName() string {
	return "user_mute"
}
