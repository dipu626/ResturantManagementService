package routes

import (
	"golang-resturant-management/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/users", controllers.GetUsers())
	incommingRoutes.GET("users/:user_id", controllers.GetUser())
	incommingRoutes.POST("/users/signup", controllers.SignUp())
	incommingRoutes.POST("/users/login", controllers.Login())

}
