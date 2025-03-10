package main

import (
	"fmt"

	"example.com/PriceCalculator/filemanager"
	"example.com/PriceCalculator/prices"
)

type IOManager interface {
	ReadLines() ([]string, error)
	WriteResult(data interface{}) error
}

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	// for range...range直後の変数に格納された値を、for文直後の第一変数分、第二変数に代入していく。
	// for文直後の1つ目の変数...ループ回数を指定
	// for文直後の2つ目の変数...配列の値が、range直後の変数に入った値が順番に代入されていく。
	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		// filemanager.New()...inputPathとoutputPathを定義
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))

		//cmdmanager.goに記載されたメソッドを呼び出せる変数cmdm（CMDManager 型のオブジェクト）を作成
		//cmdm := cmdmanager.New()

		//「税抜き価格」と「税率」を返すメソッド
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)

		//JSONファイルを作成。エラーがあれば、変数errにエラーを格納
		//go...GoRoutineを使うことにより、上記のNewTaxIncludedPriceJob()と同時並行で処理を実施
		go priceJob.Process(doneChans[index], errorChans[index])

		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// }
	}

	for index := range taxRates {
		select {
		//エラーがあれば変数errを出力
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		//正常終了なら"Done!"を出力
		case <-doneChans[index]:
			fmt.Println("Done!")
		}
	}
}
