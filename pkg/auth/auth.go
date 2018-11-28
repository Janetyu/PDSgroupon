package auth

import "golang.org/x/crypto/bcrypt"

// 用bcrypt库加密纯文本
func Encrypt(source string) (string, error) {
	// GenerateFromPassword以给定的代价返回密码的bcrypt哈希值。
	// 如果给定的成本低于MinCost，则成本将设置为DefaultCost,类似于盐值
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

// Compare compares the encrypted text with the plain text if it's the same.
func Compare(hashedPassword, password string) error {
	// CompareHashAndPassword将返回的散列密码与其明文版本进行比较
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
