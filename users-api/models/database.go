package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `db:"id"`
	Username     string    `db:"username"`
	UserHandle   string    `db:"user_handle"`
	UserBio      string    `db:"user_bio"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	Followers    uint32    `db:"followers"`
	Follows      uint32    `db:"follows"`
}

func (User) TableName() string {
	return "user"
}

type UserFollow struct {
	FollowerID uuid.UUID `db:"follower_id"`
	FollowedID uuid.UUID `db:"followed_id"`
	CreatedAt  time.Time `db:"created_at"`
}

func (UserFollow) TableName() string {
	return "user_follow"
}

type UserBlock struct {
	BlockerID uuid.UUID `db:"blocker_id"`
	BlockedID uuid.UUID `db:"blocked_id"`
	CreatedAt time.Time `db:"created_at"`
}

func (UserBlock) TableName() string {
	return "user_block"
}

type UserMute struct {
	MuterID   uuid.UUID `db:"muter_id"`
	MutedID   uuid.UUID `db:"muted_id"`
	CreatedAt time.Time `db:"created_at"`
}

func (UserMute) TableName() string {
	return "user_mute"
}
