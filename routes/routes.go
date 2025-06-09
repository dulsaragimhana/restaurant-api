package routes

import (
	"github.com/dulsaragimhana/restaurant-api/controllers"
	"github.com/dulsaragimhana/restaurant-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		restaurant := api.Group("/:restaurantID")
		restaurant.Use(middleware.AuthMiddleware())
		{
			restaurant.POST("/orders", controllers.CreateOrder)
			restaurant.GET("/orders/table/:tableNumber", controllers.GetOrderByTable)
			restaurant.GET("/orders/:orderID", controllers.GetOrderByID)
			restaurant.POST("/orders/:orderID/payment", controllers.SubmitPayment)
		}
	}
}
