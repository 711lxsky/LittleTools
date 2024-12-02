package util

import "golang.org/x/crypto/bcrypt"

func HashEncrypt(str string) (string, error) {
	// 将str转换为字节数组
	strBytes := []byte(str)
	// 直接调用bcrypt
	hashedBytes, err := bcrypt.GenerateFromPassword(strBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func CheckHashValid(passwordHashed, passwordRaw string) bool {
	// 前者是哈希之后的数据，后者是前端发来的原始密码
	err := bcrypt.CompareHashAndPassword([]byte(passwordHashed), []byte(passwordRaw))
	return err == nil
}
