package routes

import (
	"net/http"

	"github.com/1024u/go-crud/Section11_REST-API/models"
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
