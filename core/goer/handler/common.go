package handler

import (
	"github.com/gin-gonic/gin"
	"goer/config"
	MyErr "goer/error"
	"net/http"
)

func getUserIdFromContext(c *gin.Context) int {
	userId, exists := c.Get(config.TokenClaimUserId)
	if !exists {
		ResponseFail(c, http.StatusUnauthorized, MyErr.ContextError, "")
		return MyErr.IntErrValue
	}
	// 将 userId 转换为 int 类型
	userIdInt, ok := userId.(int)
	if !ok {
		ResponseFail(c, http.StatusInternalServerError, MyErr.TypeAssertionError, "")
		return MyErr.IntErrValue
	}
	return userIdInt
}
