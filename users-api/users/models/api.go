package models

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email_address"`
	Password string `json:"password"`
}

type IndividualUserRequest struct {
	ID string `json:"user_id"`
}

type FollowUserRequest struct {
	FollowerID string `json:"follower_id"`
	FollowedID string `json:"followed_id"`
}
