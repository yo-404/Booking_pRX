package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yo-404/Booking_pRX/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controllers.Signup())
	incomingRoutes.POST("users/login", controllers.Login())

}
