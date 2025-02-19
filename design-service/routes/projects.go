package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getProjects(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later"})
		return
	}
	context.JSON(http.StatusOK, events)
}
