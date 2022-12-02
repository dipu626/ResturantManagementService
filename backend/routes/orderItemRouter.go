package routes

import (
	"golang-resturant-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/orderItems", controllers.GetOrderItems())
	incommingRoutes.GET("/orderItems/:orderItem_id", controllers.GetOrderItem())
	incommingRoutes.GET("/orderItems-order/:order_id", controllers.GetOrderItemsByOrder())
	incommingRoutes.POST("/orderItems", controllers.CreateOrderItem())
	incommingRoutes.POST("/orderItems/:orderItem_id", controllers.UpdateOrderItem())

}
