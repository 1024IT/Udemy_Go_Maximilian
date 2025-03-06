// 概要：ユーザから受け付けたユーザ情報を、user.goファイルのUser structsに格納し、その情報を出力
package main

import (
	"fmt"

	"example.com/structs/user" //インポートしたいファイルのパスを記載
)

func main() {
	userfirstName := getUserData("Please enter your first name: ")
	userlastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	//概要：importしてきたuser.goファイルの中の、User stractsの情報をポインタ変数に格納
	//appUser...ポインタ変数
	//user...ポインタ型
	var appUser *user.User

	//概要：importしてきたuser.goファイルのNewUserメソッドを呼び出し、ユーザ情報を取得
	appUser, err := user.NewUser(userfirstName, userlastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	//概要：importしてきたuser.goファイルのNewAdminメソッドを呼び出し、管理者情報を取得
	admin := user.NewAdmin("text@example.com", "test123")

	admin.User.OutputUserDetails()
	admin.User.ClearUserName()
	admin.User.OutputUserDetails()

	//user structsで宣言した各変数に、値を格納し、その変数群をまとめてポインタ変数appUserに格納
	// appUser = user{
	// 	firstName: userfirstName,
	// 	lastName:  userlastName,
	// 	birthdate: userBirthdate,
	// 	createdAt: time.Now(),
	// }

	//&+ポインタ変数...ポインタ変数に格納された値がアドレスになる。
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
