package entity

import "time"

type Comment struct {
	ID             int64 `gorm:"primary_key"`
	Commenter      string
	CommentContent string
	CommentTime    time.Time
	Post           Category `gorm:"ForeignKey:PostID;AssociationForeignKey:ID"`
	PostID         int64
	Read           bool
}

func (Comment) TableName() string {
	return "bgo_post_comment"
}
