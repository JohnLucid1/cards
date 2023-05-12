package controllers

import (
	"backend/database"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllPosts(context *gin.Context) {
	posts := []models.Post{}
	record := database.Instance.Find(&posts)

	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"posts": record})
}
