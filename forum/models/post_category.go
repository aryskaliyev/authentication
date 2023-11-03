package models

type PostCategory struct {
	PostId       int    `json:"post_id,omitempty"`
	CategoryId   int    `json:"id"`
	CategoryName string `json:"name,omitempty"`
}
