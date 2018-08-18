package model

import "time"

type CommentVO struct {
	ID             int64     `json:"id"`
	Commenter      string    `json:"commenter"`
	CommentContent string    `json:"comment_content"`
	CommentTime    time.Time `json:"comment_time"`
	PostTitle      string    `json:"post_title"`
	PostID         string    `json:"post_id"`
}
