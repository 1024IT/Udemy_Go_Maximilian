package middlewares

import (
	"net/http"

	"github.com/1024u/go-crud/Section11_REST-API/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	//リクエストのヘッダに含まれているAuthorizationヘッダの値（トークン）を取得
	token := context.Request.Header.Get("Authorization")

	//未認証のリクエスト(リクエストにトークンが含まれていなければ）処理を中断し以降の処理を実施しない。
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	//取得したトークンが有効かどうかを検証
	userId, err := utils.VerifyToken(token)

	if err != nil {
		//トークンが無効であれば処理を中断し以降の処理を実施しない。
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	//変数userIdの値を"userId"に格納
	context.Set("userId", userId)
	context.Next()
}
