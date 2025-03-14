package routes

import (
	_ "github.com/pius706975/golang-test/docs"
	"github.com/pius706975/golang-test/modules"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	APIPrefix = "/"
)

// @Summary      Home Endpoint
// @Description  Shows welcome message to developers
// @Tags         Home
// @Accept       json
// @Produce      json
// @Success      200
// @Router       / [get]
func homeHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status":  200,
		"message": "Hello Go developers",
	})
}

func RouteApp(router *gin.Engine) error {
	router.GET(APIPrefix+"/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET(APIPrefix, homeHandler)

	modules.SetupRoutes(router)

	return nil
}
