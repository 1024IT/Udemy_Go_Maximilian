package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err1 := getUserInput("Revenue: ")
	expenses, err2 := getUserInput("Expenses: ")
	taxRate, err3 := getUserInput("TaxRate: ")

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		return
	}

	// outputText("expenses: ")
	// fmt.Scan(&expenses)

	// outputText("taxRate: ")
	// fmt.Scan(&taxRate)

	//ユーザーが入力した値(revenue, expenses, taxRate)を基にメソッド内で計算し、ebt, profit, ratioに格納している。
	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	//%f...フォーマット指定子の一種。浮動小数点数を10進数表記で表示する。
	//1...小数点1位まで表示する。
	//
	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.1f\n", ratio)
	storeResults(ebt, profit, ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64

	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Value must be a positive number.")
	}

	//nilの意味...エラーは無いんだけど、返却値(float64, error)に、errorが含まれているので、何か返さないといけない時に使う値
	//↑補足：エラーがあった時点で、上記のif文に記載の通りreturn 0されるので、ここのnilは無意味なコードなのだが、メソッドの返却値が（float64, error）なので
	return userInput, nil
}

func outputText(text string) {
	fmt.Println(text)
}

// 入金または出金をした際に、預金額を更新するメソッド
func storeResults(ebt, profit, ratio float64) {
	//ユーザーの入力値（ebt, profit, ratio）を変数resultsに格納
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	//テキストファイルresults.txtに記載された預金額に、変数balanceTextの値で加減
	os.WriteFile("results.txt", []byte(results), 0644)
}
