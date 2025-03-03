package main

import (
	"fmt"

	"example.com/bank/fileops"           //fileopsファイルに記載されたメソッド
	"github.com/Pallinder/go-randomdata" //go getによってgo.modファイルに記載された外部パッケージ
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------")

		//panic()...エラー発生箇所をユーザーに教える
		panic("Can't continue, sorry.")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

	for {
		presentOptions()

		//定義
		var choice int
		fmt.Print("Your choice: ")

		//入力受付
		fmt.Scan(&choice)

		//wantsCheckBakance := choice == 1

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")

			var depositAmount float64
			fmt.Scan(&depositAmount)

			//if 0円より低い金額を出金しようとした時
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				//return //←これがあると後続の処理が実行されない。
				continue //for文の最初に戻る
			}

			accountBalance += depositAmount //accountBalance = accountBalance + withdrawalAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("Your withdrawal: ")

			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			//if 0円より低い金額を出金しようとした時
			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue //for文の最初に戻る
			}

			//if 預金額を超える金額を出金しようとした時
			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				continue //for文の最初に戻る
			}

			accountBalance -= withdrawalAmount //accountBalance = accountBalance - withdrawalAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 4:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")
			return
			//break
		}
	}
}
