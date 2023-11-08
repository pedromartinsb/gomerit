package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pedromartinsb/gomerit/docs"
	"github.com/pedromartinsb/gomerit/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize Handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		// Opening
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)

		// Holding
		v1.GET("/holding", handler.ShowHoldingHandler)
		v1.POST("/holding", handler.CreateHoldingHandler)
		v1.PUT("/holding", handler.UpdateHoldingHandler)
		v1.DELETE("/holding", handler.DeleteHoldingHandler)
		v1.GET("/holdings", handler.ListHoldingsHandler)

		// Company
		v1.POST("/company", handler.CreateCompanyHandler)
	}

	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
