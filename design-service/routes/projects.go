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
