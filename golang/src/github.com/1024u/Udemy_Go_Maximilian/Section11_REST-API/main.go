package main

import (
	"github.com/1024u/Udemy_Go_Maximilian/Section11_REST-API/db"
	"github.com/1024u/Udemy_Go_Maximilian/Section11_REST-API/routes"
	"github.com/gin-gonic/gin"
)

// http://localhost:8080/eventsで、ブラウザから画面にアクセス可能
func main() {
	db.InitDB()

	//gin.Default()...HTTP(GET, POST, PUT, DELETE)を提供するメソッド。
	// 				  エンドポイントの追加やミドルウェアの登録を行う。
	server := gin.Default()

	//処理のルーティングを行うメソッド(どの処理に進むか案内してくれるメソッド)
	//※ルーティング...ネットワーク上でデータを転送する際に、その経路を決定するプロセス
	routes.RegisterRoutes(server)

	//Run()...
	server.Run(":8080") //localhost:8080
}
