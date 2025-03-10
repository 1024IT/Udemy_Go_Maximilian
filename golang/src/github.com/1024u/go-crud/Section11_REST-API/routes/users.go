package routes

import (
	"net/http"

	"github.com/1024u/go-crud/Section11_REST-API/models"
	"github.com/1024u/go-crud/Section11_REST-API/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	//.ShouldBindJSON()...リクエストボディのJSONデータをGoの構造体にバインド（マッピング）する。
	err := context.ShouldBindJSON(&user)

	if err != nil {
		//.H...mapのショートカット
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.Save()

	if err != nil {
		//.H...mapのショートカット
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		//.H...mapのショートカット
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	//入力されたログイン情報（mail、password）の組み合わせが正しいか認証
	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
		return
	}

	//トークンを生成
	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})
}
