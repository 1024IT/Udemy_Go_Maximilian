// 処理のルーティングを実施している（main.goファイルから呼び出されてる）
package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	//.GET()...エンドポイント"/events"にGETリクエストが来た時、変数getEventsによりメソッドgetEvents()が実行される
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) //:id...動的パラメータ。任意のidが「:id」に格納される
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)
	server.POST("/signup", signup)
}
