package entity

import "time"

type Post struct {
	ID         int64 `gorm:"primary_key"`
	Title      string
	Content    string
	User       User `gorm:"ForeignKey:UserID;AssociationForeignKey:ID"`
	UserID     int64
	Category   Category `gorm:"ForeignKey:CategoryID;AssociationForeignKey:ID"`
	CategoryID int64
	CreatedAt  time.Time
	IsDeleted  bool
}

func (Post) TableName() string {
	return "bgo_post"
}
