package user

import (
	"errors"
	"fmt"
	"time"
)

// ユーザから入力された情報をここで格納している。（仮想DB的な）
// structs...複数の変数をまとめて管理できるデータ構造
// 注意：別のファイルに記載したmainメソッドからこのstructを呼び出す場合、変数名は大文字にする。
type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// structs...複数の変数をまとめて管理できるデータ構造
// Admin...管理者のこと
type Admin struct {
	Email    string
	Password string
	User     User //上記のuser structsの内容を取得している
}

// ユーザ情報を出力するメソッド
// * + 変数...ポインタ変数に紐づいている値にアクセスできる。
// (u user)...レシーバ。user structsの中の値を使えるようになる。
func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

// user structsで値が格納された変数の中身をクリアするメソッド。
// userに*が付いている理由...*がないと、user structsからコピーしてきた値を編集するだけになり
//
//	元の値はclearされない。よって、*をつけてアドレスに紐づいた
//	値を変えることにより、元の値を変えている。
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// newXxx()...constructor function（コンストラクタ関数）。引数で受け取った値を元に構造体を生成する。
func NewAdmin(email, password string) Admin {
	return Admin{
		Email:    email,
		Password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthdate: "---",
			createdAt: time.Now(),
		},
	}
}

// newXxx()...constructor function（コンストラクタ関数）。引数で受け取った値を元に構造体を生成する。
//
//	　ここでは、値を受け取った後($user)、pointer(*user)で返す処理を行なっている
func NewUser(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First name, last name and birthdate are required.")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
