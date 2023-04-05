package model

import "time"

type Book struct {
	BookID     uint64 `json:"id" gorm:"column:id;primaryKey;autoIncrement;not null"`
	Title      string `json:"title" gorm:"column:title;unique;not null;type:varchar(191)"`
	Author     string `json:"author" gorm:"column:author"`
	Desc       string `json:"desc" gorm:"column:desc"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
