package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"password"`
	Videos    []Video   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"videos"`
	Comments  []Comment `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"comments"`
	CreatedAt time.Time `json:"createdAt"`
}

type Video struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Views       int       `json:"views"`
	Likes       int       `json:"likes"`
	User        User      `json:"user"`
	UserID      int       `json:"userId"`
	Comments    []Comment `gorm:"foreignKey:VideoID;constraint:OnDelete:CASCADE;" json:"comments"`
	CreatedAt   time.Time `json:"createdAt"`
}

type Comment struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Content   string    `gorm:"type:text"  json:"content"`
	Likes     int       `json:"likes"`
	Dislikes  int       `json:"dislikes"`
	User      User      `json:"user"`
	UserID    int       `json:"userId"`
	Video     Video     `json:"video"`
	VideoID   int       `json:"videoId"`
	CreatedAt time.Time `json:"createdAt"`
}
