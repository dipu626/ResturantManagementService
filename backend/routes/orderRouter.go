package routes

import (
	"golang-resturant-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/orders", controllers.GetOrders())
	incommingRoutes.GET("/orders/:order_id", controllers.GetOrder())
	incommingRoutes.POST("/orders", controllers.CreateOrder())
	incommingRoutes.POST("/orders/:order_id", controllers.UpdateOrder())
}
