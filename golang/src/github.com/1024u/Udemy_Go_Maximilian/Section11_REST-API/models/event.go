// イベントに対するCRUDが記載
package models

import (
	"fmt"
	"time"

	"github.com/1024u/Udemy_Go_Maximilian/Section11_REST-API/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

// スライスの宣言
var events = []Event{}

func (e *Event) Save() error {
	//?...プレースホルダー。SQLiteで使用される特別な記号で、クエリ実行時にデータが渡される場所
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)`

	//変数queryに格納されたクエリを、変数stmtに格納
	//Prepare()...指定したクエリをコンパイルすることにより、バグを検知するメソッド
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	//defer Close()...Save()の処理が終わったらステートメントを閉じる
	defer stmt.Close()

	fmt.Println(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	//Exec()...データの変更(INSERT, UPDATE, DELETE)を行うクエリを実行する。
	// 		   具体的には、変数stmtに格納されたクエリの「?(プレースホルダ)」に、Event structで取得したデータを格納
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return err
	}

	//LastInsertId()...INSERTクエリによって新しく挿入された行のIDを取得
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	//Query()...データの取得(SELECT)を行うクエリを実行する。
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	//defer Close()...Save()の処理が終わったらステートメントを閉じる
	defer rows.Close()

	var events []Event

	//.Next()...クエリの実行結果の次の行が存在するかチェックするメソッド
	for rows.Next() {
		var event Event

		//クエリの実行結果の各行の値をeventオブジェクトに格納
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		//eventオブジェクトに格納された値を、スライスeventsに追加
		events = append(events, event)
	}

	return events, nil
}

// IDから該当のイベント情報を取得→イベント情報をEvent structに格納
func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"

	//QueryRow()...指定したidに一致する1行のデータを取得するメソッド
	//			   第一引数...実行するクエリを記載
	//			　　第二引数...第一引数で指定したクエリの「？」に入れる数字
	row := db.DB.QueryRow(query, id)

	var event Event

	//クエリの実行結果の各行の値をeventオブジェクトに格納
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}

	//指定したIDのイベント情報を返す
	return &event, nil
}

func (event Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	//Exec()...データの変更(INSERT, UPDATE, DELETE)を行うクエリを実行する。
	// 		   具体的には、変数stmtに格納されたクエリの「?(プレースホルダ)」に、Event structで取得したデータを格納
	_, err = stmt.Exec(event.ID)

	return err
}

func (e Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}

func (e Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}
