package model

type Tags struct {
	ID   int    `gorm:"primary_key"`
	Name string `gorm:"size:255"`
}
