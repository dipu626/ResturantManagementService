package routes

import (
	"golang-resturant-management/controllers"

	"github.com/gin-gonic/gin"
)

func NoteRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/notes", controllers.GetNotes())
	incommingRoutes.GET("/notes/:note_id", controllers.GetNote())
	incommingRoutes.POST("/notes", controllers.CreateNote())
	incommingRoutes.POST("/notes/:note_id", controllers.UpdateNote())
}
