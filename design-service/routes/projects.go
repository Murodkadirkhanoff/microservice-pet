package routes

import (
	"net/http"

	"github.com/Murodkadirkhanoff/uiux-design-service/models"
	"github.com/gin-gonic/gin"
)

func getProjects(context *gin.Context) {
	designs, err := models.GetAllProjects()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later"})
		return
	}
	context.JSON(http.StatusOK, designs)
}

func CreateProject(context *gin.Context) {
	var design models.Design
	err := context.ShouldBindJSON(&design)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	userID := 1
	design.UserID = int64(userID)
	err = design.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create design.", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Design Created successfully", "design": design})
}
