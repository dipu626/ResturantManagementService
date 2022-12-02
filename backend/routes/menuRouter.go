package routes

import (
	"golang-resturant-management/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/menus", controllers.GetMunus())
	incommingRoutes.GET("/menus/:menu_id", controllers.GetMenu())
	incommingRoutes.POST("/menus", controllers.CreateMenu())
	incommingRoutes.POST("/menus/:menu_id", controllers.UpdateMenu)
}
