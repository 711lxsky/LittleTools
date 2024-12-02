package util

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"goer/config"
	MyErr "goer/error"
	"time"
)

func GenerateTokenWithUserInfo(userId int) (string, error) {
	// 设置过期时间
	expireAt := time.Now().Add(time.Duration(config.TokenExpireTimeDays) * 24 * time.Hour)
	claims := jwt.MapClaims{}
	claims[config.TokenClaimAuthorized] = true
	claims[config.TokenClaimAuthority] = config.TokenClaimAuthorityValue
	claims[config.TokenClaimUserId] = userId
	claims[config.TokenClaimExpireTime] = expireAt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(config.TokenSecret))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func CheckTokenValid(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(config.TokenSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims[config.TokenClaimExpireTime].(float64) {
			return nil, errors.New(MyErr.TokenExpired)
		}
		return token, nil
	}
	return nil, errors.New(MyErr.TokenInvalid)
}

func GetUserInfoFromToken(tokenString string) (int, error) {
	token, err := CheckTokenValid(tokenString)
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New(MyErr.TokenClaimsError)
	}

	userId, ok := claims[config.TokenClaimUserId].(float64)
	if !ok {
		return 0, errors.New(MyErr.TokenClaimsError)
	}

	return int(userId), nil
}
