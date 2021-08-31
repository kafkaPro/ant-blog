package utils

import "golang.org/x/crypto/bcrypt"

// CompareHashAndPassword
// 并没有对密码进行加盐处理(salt)，而是直接比较hash后的密码和加密存储的密码是否相等
// 主要是验证密码阶段使用
func CompareHashAndPassword(p1 string, p2 string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(p1), []byte(p2)); err != nil {
		return false
	}
	return true
}

// EncryptPassword
// 对用户的密码进行加密，加密后的密码会存储在数据库当中
func EncryptPassword(pwd string) (epwd string, err error) {
	if bs, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost); err != nil {
		return string(bs), nil
	}
	return "", err
}
