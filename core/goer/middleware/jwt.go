package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"goer/config"
	myErr "goer/error"
	"goer/handler"
	"goer/util"
	"net/http"
	"strings"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(config.TokenName)
		if authHeader == "" {
			handler.ResponseFail(c, http.StatusUnauthorized, myErr.TokenAuthorityHeaderMissed, "")
			return
		}
		tokenStr := strings.TrimPrefix(authHeader, config.TokenHeader)
		if tokenStr == authHeader || tokenStr == "" {
			handler.ResponseFail(c, http.StatusUnauthorized, myErr.TokenAuthorityHeaderMissed, "")
			return
		}
		token, err := util.CheckTokenValid(tokenStr)
		if err != nil {
			handler.ResponseFail(c, http.StatusUnauthorized, err.Error(), "")
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// 将用户信息存储在上下文中，方便后续调用
			c.Set(config.TokenClaimUserId, claims[config.TokenClaimUserId])
		}
		c.Next()
	}
}
