package hash

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// 随机字符串
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// 使用当前时间作为seed
var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// 创建Salt
func CreateSalt() string {
	b := make([]byte, bcrypt.MaxCost)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
