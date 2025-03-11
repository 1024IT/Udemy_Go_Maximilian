// イベントのCRUDを呼び出している。
package routes

import (
	"net/http"
	"strconv"

	"github.com/1024u/Udemy_Go_Maximilian/Section11_REST-API/models"
	"github.com/gin-gonic/gin"
)

// GETリクエストの処理（データ取得）（ヒットした全てのイベントを取得する）
func getEvents(context *gin.Context) {
	//.GetAllEvents()...event.goのEvent structのデータを全取得
	events, err := models.GetAllEvents()

	if err != nil {
		//.H...mapのショートカット
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}
	//.JSON()...レスポンスとして JSON を返すメソッド。
	//			第一引数...HTTPステータスコードを指定（200, 400, 401など）
	//			第二引数...JSONとして変換したいGoのデータを指定
	context.JSON(http.StatusOK, events)
}

// GETリクエストの処理（データ取得）（idで指定した特定のイベントを取得する）
func getEvent(context *gin.Context) {
	//Param()...パスパラメータを取得するためのメソッド。
	//			第一引数...変換したい値
	//			第二引数...進数
	//			第三引数...ビット
	//PaserInt()...StringからIntに変換
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}

	//GetEventByID...event.goで定義されている、IDからイベントを取得するメソッド
	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
	}

	context.JSON(http.StatusOK, event)
}

// POSTリクエストの処理（データ作成）
func createEvent(context *gin.Context) {

	//models.Event...event.goのEvent structの空の変数を宣言
	var event models.Event
	//.ShouldBindJSON()...リクエストボディのJSONデータをGoの構造体にバインド（マッピング）する。
	err := context.ShouldBindJSON(&event)

	if err != nil {
		//.H...mapのショートカット
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	//"userId"の値を取得
	userId := context.GetInt64("userId")

	//Event structに、イベントを作成したユーザのIDを格納
	event.UserID = userId

	//event.goのEvent structに値を格納(セーブする)
	event.Save()

	// if err != nil {
	// 	context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create events. Try again later."})
	// 	return
	// }

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func updateEvent(context *gin.Context) {
	//Param()...パスパラメータを取得するためのメソッド。
	//			第一引数...変換したい値
	//			第二引数...進数
	//			第三引数...ビット
	//PaserInt()...StringからIntに変換
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}

	//"userId"の値を取得
	userId := context.GetInt64("userId")

	//入力されたIDのイベント（更新前のイベント）を取得
	//※GetEventByID()の返却値はmodelsフォルダのevent.goに記載のEvent structに格納されるので、ここでデータは受け取らない。
	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event."})
		return
	}

	//指定したイベントのユーザーのIDと、ログインしたユーザーのIDが一致しているか確認
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event."})
		return
	}

	//更新処理
	//※上記のGetEventByID()によって、指定したイベントの情報が格納されたEvent structから値を取得
	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}

	//Event structのIDに、指定されたいIDを格納
	updatedEvent.ID = eventId
	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully!"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}

	//ログインしたユーザーのID("userId")の値を取得
	userId := context.GetInt64("userId")

	//入力されたIDのイベント（更新前のイベント）を取得
	//※GetEventByID()の返却値はmodelsフォルダのevent.goに記載のEvent structに格納されるので、ここでデータは受け取らない。
	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event."})
		return
	}

	//指定したイベントのユーザーのIDと、ログインしたユーザーのIDが一致しているか確認
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event."})
		return
	}

	//DELETE文を実行する処理を呼び出す
	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully!"})
}
