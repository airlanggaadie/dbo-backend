package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h Handler) authMiddleware(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	splitToken := strings.Split(token, " ")
	if splitToken[0] != "Bearer" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "invalid authorization token",
		})
		return
	}

	userId, err := h.jwtRepository.Verify(splitToken[1])
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.Set("user_id", userId)

	ctx.Next()
}
