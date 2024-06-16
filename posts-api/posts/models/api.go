package models

type CreatePostRequest struct {
	ParentID *string `json:"parent_post_id"`
	UserID   string  `json:"user_id"`
	Contents string  `json:"contents"`
}

type IndividualPostRequest struct {
	ID string `json:"post_id"`
}

type GetFeedRequest struct {
	ID string `json:"user_id"`
}
