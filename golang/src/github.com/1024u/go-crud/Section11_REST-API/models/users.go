package models

import (
	"errors"

	"github.com/1024u/go-crud/Section11_REST-API/db"
	"github.com/1024u/go-crud/Section11_REST-API/utils"
)

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

	//パスワードをハッシュ化（元のデータを不規則な文字列（ハッシュ値）に変換）
	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	//LastInsertId()...INSERTクエリによって新しく挿入された行のIDを取得
	userId, err := result.LastInsertId()

	u.ID = userId
	return err
}

func (u *User) ValidateCredentials() error {
	//メールアドレスをキーに登録済みのパスワードを取得
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	//.Scan()...変数rowに格納された値(パスワード)をretrievedPasswordに格納
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return err
	}

	//既にハッシュ化されているパスワードと、入力パスワードを比較
	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("Credentials invalid")
	}

	return nil

}
