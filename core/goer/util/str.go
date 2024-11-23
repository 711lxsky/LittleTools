package util

import (
	"crypto/rand"
	"math/big"
)

// GenerateRandomString 生成指定长度的随机字符串
func GenerateRandomString(length int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
		return ""
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes)
}

// GenerateRandomStringWithRangeLength 生成指定长度范围内的随机字符串
func GenerateRandomStringWithRangeLength(minLength, maxLength int) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length, err := rand.Int(rand.Reader, big.NewInt(int64(maxLength-minLength+1)))
	if err != nil {
		return "", err
	}
	length = length.Add(length, big.NewInt(int64(minLength)))

	bytes := make([]byte, length.Int64())
	for i := range bytes {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		bytes[i] = letters[num.Int64()]
	}
	return string(bytes), nil
}
