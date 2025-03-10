package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	//string→byte
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// 既にハッシュ化されているパスワードと、入力パスワードを比較
func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	//エラーが無ければreturn
	return err == nil
}
