// 処理のルーティングを実施している（main.goファイルから呼び出されてる）
package routes

import (
	"github.com/1024u/go-crud/Section11_REST-API/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	//.GET()...エンドポイント"/events"にGETリクエストが来た時、変数getEventsによりメソッドgetEvents()が実行される
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) //:id...動的パラメータ。任意のidが「:id」に格納される

	//server.Group("/") を使い、authenticatedというルートグループを作成
	//このグループに登録したルートはすべて共通の設定を持つ。
	authenticated := server.Group("/")

	//authenticatedグループに属したすべてのルートにミドルウェア（認証）を実施
	authenticated.Use(middlewares.Authenticate)

	//各ルートをまとめて定義
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
