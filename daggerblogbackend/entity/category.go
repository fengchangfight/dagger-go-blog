package entity

type Category struct {
	ID   int64 `gorm:"primary_key"`
	Name string
}

func (Category) TableName() string {
	return "bgo_category"
}
