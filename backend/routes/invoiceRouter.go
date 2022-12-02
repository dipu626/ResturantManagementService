package routes

import (
	"golang-resturant-management/controllers"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/invoices", controllers.GetInvoices())
	incommingRoutes.GET("/invoices/:invoice_id", controllers.GetInvoice())
	incommingRoutes.POST("/invoices", controllers.CreateInvoices())
	incommingRoutes.POST("/invoices/:invoce_id", controllers.UpdateInvoice())
}
