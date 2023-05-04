package models

import (
	"time"
)

type Game struct {
	ID        int
	Name      string
	Genre     string
	Developer string
	Year      int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Category struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Post struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	Title        string    `gorm:"not null"`
	Description  string    `gorm:"not null"`
	Photo        string    `gorm:"not null"`
	DateReleased time.Time `gorm:"not null"`
	CategoryID   uint      `gorm:"not null"`
	Category     Category  `gorm:"foreignKey:CategoryID"`
	Reviews      []Review  `gorm:"foreignKey:PostID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Review struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Star      uint      `gorm:"not null"`
	PostID    uint      `gorm:"not null"`
	Post      Post      `gorm:"foreignKey:PostID"`
	Comments  []Comment `gorm:"foreignKey:ReviewID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Comment struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"not null"`
	Subject   string `gorm:"not null"`
	Message   string `gorm:"not null"`
	ReviewID  uint   `gorm:"not null"`
	Review    Review `gorm:"foreignKey:ReviewID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
