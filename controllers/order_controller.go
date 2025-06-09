package controllers

import (
	"net/http"

	"github.com/dulsaragimhana/restaurant-api/models"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var req models.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Call Square API and persist to DB
	c.JSON(http.StatusOK, gin.H{"message": "Order created"})
}

func GetOrderByTable(c *gin.Context) {
	table := c.Param("tableNumber")
	// TODO: Fetch from DB
	c.JSON(http.StatusOK, gin.H{"table": table})
}

func GetOrderByID(c *gin.Context) {
	orderID := c.Param("orderID")
	// TODO: Fetch from DB
	c.JSON(http.StatusOK, gin.H{"orderID": orderID})
}

func SubmitPayment(c *gin.Context) {
	orderID := c.Param("orderID")
	var req models.PaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// TODO: Submit payment to Square and store
	c.JSON(http.StatusOK, gin.H{"message": "Payment submitted for order: " + orderID})
}
