package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWT の署名に使う秘密鍵。JWT の改ざんを防ぐために、この鍵を使ってハッシュ化する。
const secretKey = "supersecret"

func GenerateToken(email string, userId int64) (string, error) {
	//新しいJWT(JSON Web Token)トークンを作成
	//第一引数...JWTの署名アルゴリズム
	//第二引数...JWTに含める、ユーザのログインデータ
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		//2時間でトークンが失効期限する
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	//改ざん防止のため、secretKeyを使ってJWTに署名→署名付きJWTを返す
	return token.SignedString([]byte(secretKey))
}

// トークンの有効性を確認
func VerifyToken(token string) (int64, error) {
	//Parse()...トークンの署名アルゴリズム を確認し、署名の検証 が行える準備する。
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

		//トークンの署名アルゴリズムが HMAC（HS256 など） であることを確認
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Unexpected signing method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("Could not parse token.")
	}

	//実際にトークンが有効かどうか（署名が正しく、期限切れでないか）を確認
	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, errors.New("Invalid token!")
	}

	claims := parsedToken.Claims.(jwt.MapClaims)

	// if !ok {
	// 	return errors.New("Invalid token claims")
	// }

	// email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))
	return userId, nil
}
