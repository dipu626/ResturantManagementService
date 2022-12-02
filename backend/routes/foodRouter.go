package routes

import (
	"golang-resturant-management/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/foods", controllers.GetFoods())
	incommingRoutes.GET("/foods/:food_id", controllers.GetFood())
	incommingRoutes.POST("/foods", controllers.CreateFood())
	incommingRoutes.POST("/foods/:food_id", controllers.UpdateFood())
}
