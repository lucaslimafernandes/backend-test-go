package services

import (
	"backendtest-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MarkUnsafeAPI(c *gin.Context) {

	var fileReviewInput models.FileReviewInput

	err := c.ShouldBindJSON(&fileReviewInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review := models.FileReview{
		FileID:     fileReviewInput.FileID,
		ReviewerId: fileReviewInput.ReviewerId,
		Unsafe:     fileReviewInput.Unsafe,
	}

	models.DB.Create(&review)

	c.JSON(http.StatusOK, gin.H{"response": review})

}
