package main

import (
	"fmt"
	//"Udemy_Go_Maximilian/initializers"
	"math"
)

// 定義
const inflationRate = 2.5

// func init() {
// 	initializers.LoadEnvVariables()
// 	initializers.ConnectToDB()
// }

func main() {
	//定義
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	//入力受付
	//Scan...コマンドプロンプトからのインプットを受け入れる
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	outputText("Years: ")
	fmt.Scan(&years)

	//計算
	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	//出力
	//fmt.Println("Future Value:", futureValue)
	//%v..."fmt"パッケージにおけるフォーマット指定子の一種。構造体（複数の値を格納した変数的なもの）を出力する際、フィールド名が表示される(https://pkg.go.dev/fmt)
	//%f..."fmt"パッケージにおけるフォーマット指定子の一種。浮動小数点数を10進数表記で表示する
	//.0...小数点を0とする。
	//\n...フォーマット指定子の一種。改行する
	//Sprintf...任意の型と文字列をまとめて文字列にする（それを変数に格納し、その変数を出力することが多い）（https://qiita.com/Sekky0905/items/5a65602dce83551184b3）
	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value(adjusted for inflation): %.1f", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

// メソッド名(取得値の型)（返却値の型）
func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	//.Pow()...べき乗計算を行う関数
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}

// func main() {
// 	r := gin.Default()
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "pingpongpang",
// 		})
// 	})
// 	r.Run()
// }
