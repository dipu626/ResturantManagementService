package controllers

import (
	"context"
	"fmt"
	"golang-resturant-management/database"
	"golang-resturant-management/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderItemPack struct {
	Table_id    *string
	Order_items []models.OrderItem
}

var orderItemCollection *mongo.Collection = database.OpenCollection(database.Client, "orderItem")

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		result, err := orderItemCollection.Find(
			ctx,
			bson.M{},
		)
		defer cancel()

		if err != nil {
			msg := fmt.Sprintf("message: error occured while listening order items")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		var allOrderItems []bson.M

		if err := result.All(ctx, &allOrderItems); err != nil {
			log.Fatal(err)
			return
		}

		c.JSON(http.StatusOK, allOrderItems)
	}
}

func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var orderItem models.OrderItem
		orderItemId := c.Param("order_item_id")

		err := orderItemCollection.FindOne(
			ctx,
			bson.M{
				"order_item_id": orderItemId,
			},
		).Decode(&orderItem)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listening order item by order item id"})
			return
		}
		c.JSON(http.StatusOK, orderItem)
	}
}

func GetOrderItemsByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var orderId = c.Param("order_id")

		allOrderItems, err := ItemsByOrder(orderId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listening order items by order id"})
			return
		}
		c.JSON(http.StatusOK, allOrderItems)
	}
}

func ItemsByOrder(id string) (OrderItems []primitive.M, err error) {

}

func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
