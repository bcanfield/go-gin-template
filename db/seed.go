package db

import (
	"fmt"
	"go-template/models"
	"log"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	fmt.Println("Seeding!!")
	seedUsers(db)
	seedVideos(db)
	seedComments(db)
}

func seedUsers(db *gorm.DB) {
	names := []string{"John Doe", "Jane Smith", "Michael Johnson", "Emily Davis", "David Wilson"}
	for i := 0; i < 20; i++ {
		user := models.User{
			Name:     names[rand.Intn(len(names))],
			Email:    fmt.Sprintf("user%d@example.com", i+1),
			Password: "password",
		}
		if err := db.FirstOrCreate(&user).Error; err != nil {
			log.Printf("Failed to seed user: %v", err)
		}
	}
}

func seedVideos(db *gorm.DB) {
	titles := []string{"Funny Cats Compilation", "Cooking Tutorial: Spaghetti Carbonara", "Travel Vlog: Exploring Bali", "Gaming Highlights: Fortnite Battle Royale", "Music Video: Top Hits of 2021"}
	users := getUsers(db)
	for i := 0; i < 20; i++ {
		video := models.Video{
			Title:       titles[rand.Intn(len(titles))],
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			Views:       rand.Intn(10000),
			Likes:       rand.Intn(1000),
			UserID:      int(users[rand.Intn(len(users))].ID),
		}
		if err := db.FirstOrCreate(&video).Error; err != nil {
			log.Printf("Failed to seed video: %v", err)
		}
	}
}

func seedComments(db *gorm.DB) {
	contents := []string{"Great video!", "I learned a lot from this.", "Can't stop watching this!", "Nice work!", "Keep it up!"}
	users := getUsers(db)
	videos := getVideos(db)
	for i := 0; i < 20; i++ {
		comment := models.Comment{
			Content:   contents[rand.Intn(len(contents))],
			Likes:     rand.Intn(100),
			Dislikes:  rand.Intn(50),
			UserID:    int(users[rand.Intn(len(users))].ID),
			VideoID:   int(videos[rand.Intn(len(videos))].ID),
			CreatedAt: time.Now().AddDate(0, -1, i),
		}
		if err := db.FirstOrCreate(&comment).Error; err != nil {
			log.Printf("Failed to seed comment: %v", err)
		}
	}
}

func getUsers(db *gorm.DB) []models.User {
	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		log.Fatalf("Failed to retrieve users: %v", err)
	}
	return users
}

func getVideos(db *gorm.DB) []models.Video {
	var videos []models.Video
	if err := db.Find(&videos).Error; err != nil {
		log.Fatalf("Failed to retrieve videos: %v", err)
	}
	return videos
}
