package controllers


import (
	"backend/database"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePost(context *gin.Context) {
	var post models.Post
	if err := context.ShouldBindJSON(&post); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()} )
		context.Abort()
		return
	}

	record := database.Instance.Create(&post)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, gin.H{"postId": post.ID, "question": post.Question, "answer": post.Answer})
}
