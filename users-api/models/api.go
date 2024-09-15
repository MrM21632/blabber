package models

import "github.com/google/uuid"

type CreateUserRequest struct {
	Username string  `json:"username"`
	Handle   string  `json:"handle"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Bio      *string `json:"user_bio"`
}

type IndividualUserRequest struct {
	ID uuid.UUID `json:"user_id"`
}

type UpdateUserRequest struct {
	ID       uuid.UUID `json:"user_id"`
	Username *string   `json:"username"`
	Handle   *string   `json:"handle"`
	Email    *string   `json:"email"`
	Bio      *string   `json:"user_bio"`
}

type UpdateUserPasswordRequest struct {
	ID          uuid.UUID `json:"user_id"`
	OldPassword string    `json:"old_password"`
	NewPassword string    `json:"new_password"`
}

type FollowersRequest struct {
	FollowerID uuid.UUID `json:"follower_id"`
	FollowedID uuid.UUID `json:"followed_id"`
}

type BlocksRequest struct {
	BlockerID uuid.UUID `json:"blocker_id"`
	BlockedID uuid.UUID `json:"blocked_id"`
}

type MutesRequest struct {
	MuterID uuid.UUID `json:"muter_id"`
	MutedID uuid.UUID `json:"muted_id"`
}

type PartialUser struct {
	ID       uuid.UUID `json:"user_id"`
	Handle   string    `json:"handle"`
	Username string    `json:"username"`
}
