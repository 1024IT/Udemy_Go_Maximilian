package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// mainメソッド開始時（ユーザーの取引時）に、現預金額をbalance.txtから取得するメソッド
func GetFloatFromFile(fileName string) (float64, error) {
	//テキストファイルbalance.txtに記載された値を変数dataに格納
	//ReadFile()...balance.txtを読み込み、「ファイルに記載された値(data)」と、(あれば)エラー(err)を返す。
	data, err := os.ReadFile(fileName)

	//if エラーがnil(無し)ではなければ
	if err != nil {
		return 1000, errors.New("Failed to find file.")
	}

	//変数dataの値をstring型に変換し、balanceTextに格納
	valueText := string(data)
	//変数balanceTextの値をstring型に変換。もしエラーが発生したら返す（→変数errに格納）
	value, err := strconv.ParseFloat(valueText, 64)

	//if エラーがnil(無し)ではなければ
	if err != nil {
		return 1000, errors.New("Failed to parse stored value.")
	}

	//エラーがない
	return value, nil
}

// 入金または出金をした際に、預金額を更新するメソッド
func WriteFloatToFile(value float64, fileName string) {
	//ユーザーの入力値（balance）を変数balanceTextに格納
	valueText := fmt.Sprint(value)
	//テキストファイルbalance.txtに記載された預金額に、変数balanceTextの値で加減
	os.WriteFile(fileName, []byte(valueText), 0644)
}
