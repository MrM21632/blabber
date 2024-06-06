package models

type CreateUserRequest struct {
	Handle   string `json:"handle"`
	Username string `json:"username"`
	Email    string `json:"email_address"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	ID       string  `json:"user_id"`
	Handle   *string `json:"handle"`
	Username *string `json:"username"`
	Email    *string `json:"email_address"`
}

type UpdateUserPasswordRequest struct {
	ID       string `json:"user_id"`
	Password string `json:"password"`
}

type IndividualUserRequest struct {
	ID string `json:"user_id"`
}

type FollowUserRequest struct {
	FollowerID string `json:"follower_id"`
	FollowedID string `json:"followed_id"`
}

type GetFollowersRequest struct {
	FollowedID string `json:"followed_id"`
}

type GetFollowsRequest struct {
	FollowerID string `json:"follower_id"`
}

type UserFollowerEntity struct {
	ID       string `json:"id"`
	Handle   string `json:"handle"`
	Username string `json:"username"`
}
