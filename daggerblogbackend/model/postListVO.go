package model

type PostListVO struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Category  string `json:"category"`
	CreatedAt string `json:"created_at"`
}
