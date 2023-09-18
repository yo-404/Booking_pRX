package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yo-404/Booking_pRX/controllers"
	"github.com/yo-404/Booking_pRX/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUsers())
}
