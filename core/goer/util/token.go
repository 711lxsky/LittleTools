package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"goer/config"
	MyErr "goer/error"
	"reflect"
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
		expireTimeInterface, exists := claims[config.TokenClaimExpireTime]
		if !exists {
			return nil, errors.New(MyErr.TokenInvalid)
		}
		// 添加调试信息
		fmt.Printf("expireTimeInterface type: %v, value: %v\n", reflect.TypeOf(expireTimeInterface), expireTimeInterface)
		// 尝试将 expireTimeInterface 转换为 time.Time 类型
		var expireTime time.Time
		switch v := expireTimeInterface.(type) {
		case time.Time:
			expireTime = v
		case float64:
			expireTime = time.Unix(int64(v), 0)
		case json.Number:
			intValue, err := v.Int64()
			if err != nil {
				return nil, errors.New(MyErr.TokenInvalid)
			}
			expireTime = time.Unix(intValue, 0)
		case string:
			// 使用 time.RFC3339Nano 布局字符串解析字符串
			expireTime, err = time.Parse(time.RFC3339Nano, v)
			if err != nil {
				return nil, errors.New(MyErr.TokenInvalid)
			}
		default:
			return nil, errors.New(MyErr.TokenInvalid)
		}
		if time.Now().After(expireTime) {
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
