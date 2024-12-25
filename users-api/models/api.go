package models

import "github.com/google/uuid"

type CreateUserRequest struct {
	Username string  `json:"username"`
	Handle   string  `json:"handle"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Bio      *string `json:"user_bio,omitempty"`
}

type IndividualUserRequest struct {
	ID uuid.NullUUID `json:"user_id"`
}

type UpdateUserRequest struct {
	ID       uuid.NullUUID `json:"user_id"`
	Username *string       `json:"username,omitempty"`
	Handle   *string       `json:"handle,omitempty"`
	Email    *string       `json:"email,omitempty"`
	Bio      *string       `json:"user_bio,omitempty"`
}

type UpdateUserPasswordRequest struct {
	ID          uuid.NullUUID `json:"user_id"`
	OldPassword string        `json:"old_password"`
	NewPassword string        `json:"new_password"`
}

type FollowersRequest struct {
	FollowerID uuid.NullUUID `json:"follower_id"`
	FollowedID uuid.NullUUID `json:"followed_id"`
}

type BlocksRequest struct {
	BlockerID uuid.NullUUID `json:"blocker_id"`
	BlockedID uuid.NullUUID `json:"blocked_id"`
}

type MutesRequest struct {
	MuterID uuid.NullUUID `json:"muter_id"`
	MutedID uuid.NullUUID `json:"muted_id"`
}

type PartialUser struct {
	ID       uuid.NullUUID `db:"id" json:"user_id"`
	Handle   string        `db:"user_handle" json:"handle"`
	Username string        `db:"username" json:"username"`
}
