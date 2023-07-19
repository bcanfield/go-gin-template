package handlers

import (
	"go-template/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateVideo(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var video models.Video
	if err := c.ShouldBindJSON(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	video.CreatedAt = time.Now()

	if err := db.Create(&video).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, video)
}

func GetVideo(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var video models.Video
	videoID := c.Param("id")

	if err := db.First(&video, videoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
		return
	}

	c.JSON(http.StatusOK, video)
}

func UpdateVideo(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var video models.Video
	videoID := c.Param("id")

	if err := db.First(&video, videoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
		return
	}

	if err := c.ShouldBindJSON(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&video).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, video)
}

func DeleteVideo(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var video models.Video
	videoID := c.Param("id")

	if err := db.First(&video, videoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
		return
	}

	if err := db.Delete(&video).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Video deleted"})
}

func GetVideos(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var videos []models.Video
	if err := db.Find(&videos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, videos)

}
