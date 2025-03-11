// DBの初期化とCREATE TABLEが記載
package db

import (
	"database/sql"
	// "github.com/1024u/go-crud/Section11_REST-API/models"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	//.Open()...SQLを開く
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	//アプリケーションがDBに対して同時に開ける最大の接続数を指定。
	DB.SetMaxOpenConns(10)

	//アプリケーションがDBにアイドル状態（使用されていない状態）の接続の最大数を指定。
	DB.SetMaxIdleConns(5)

	createTables()
}

// create tableを実行
func createTables() {
	// usersテーブルが既に存在している場合は削除
	dropTable := "DROP TABLE IF EXISTS users;"
	_, err := DB.Exec(dropTable)
	if err != nil {
		panic("Could not drop users table.")
	}

	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);
	`

	_, err = DB.Exec(createUserTable)

	if err != nil {
		panic("Could not create users table.")
	}

	// eventsテーブルが既に存在している場合は削除
	dropTable = "DROP TABLE IF EXISTS events;"
	_, err = DB.Exec(dropTable)
	if err != nil {
		panic("Could not drop events table.")
	}

	// テーブルeventsを生成
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);
	`
	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table.")
	}

	// registrationsテーブルが既に存在している場合は削除
	dropTable = "DROP TABLE IF EXISTS registrations;"
	_, err = DB.Exec(dropTable)
	if err != nil {
		panic("Could not drop registrations table.")
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER, 
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		panic("Could not create registrations table")
	}
}
