package model

type PostVO struct {
	ID         int64  `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Category   string `json:"category"`
	CategoryID int64  `json:"category_id"`
	CreatedAt  string `json:"created_at"`
}
