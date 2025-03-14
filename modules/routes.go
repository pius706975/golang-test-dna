package modules

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/language", CreateLanguageController)
	router.DELETE("/language/:id", DeleteLanguageController)
	router.PATCH("/language/:id", UpdateLanguageController)
	router.GET("/languages", GetLanguagesController)
	router.GET("/language/:id", GetLanguageByIDController)
	router.GET("/palindrome", GetPalindromeController)
}
