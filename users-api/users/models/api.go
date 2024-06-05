package models

type CreateUserRequest struct {
	Handle   string `json:"handle"`
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

type UserFollowerEntity struct {
	ID       string `json:"id"`
	Handle   string `json:"handle"`
	Username string `json:"username"`
}
