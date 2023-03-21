package middlewares

import (
	"gin-starter-template/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		utils.DLog().Msg("AuthMiddleware called")
		// Do some stuff here
		c.Next()
	}
}
