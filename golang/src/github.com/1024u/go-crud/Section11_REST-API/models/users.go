package models

import "github.com/1024u/go-crud/Section11_REST-API/db"

type User struct {
	ID       int64
	Email    string `binding: "required"`
	Password string `binding: "required"`
}

// ユーザー情報をusersテーブルに登録
func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(u.Email, u.Password)

	if err != nil {
		return err
	}

	//LastInsertId()...INSERTクエリによって新しく挿入された行のIDを取得
	userId, err := result.LastInsertId()

	u.ID = userId
	return err
}
