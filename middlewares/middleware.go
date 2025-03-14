package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"slices"
)

func NotFoundHandler(ctx *gin.Context) {
	methods := []string{"GET", "POST", "PUT", "DELETE", "PATCH"}
	if slices.Contains(methods, ctx.Request.Method) {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{
			"status":  405,
			"message": "Method Not Allowed",
		})
		return
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Page Not Found",
	})
}
