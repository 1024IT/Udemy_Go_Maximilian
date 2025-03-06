package prices

import (
	"fmt"

	"example.com/PriceCalculator/conversion.go"
	"example.com/PriceCalculator/iomanager"
)

type TaxIncludedPriceJob struct {
	//"-"...キーを除外
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_rate"`
}

func (job *TaxIncludedPriceJob) LoadData() error {

	//prices.txtに記載された値を変数linesに格納(filemanager.go参照)
	lines, err := job.IOManager.ReadLines()

	if err != nil {
		return err
	}

	//"prices.txt"の値をStrings→Floatsに変換(conversion.go参照)
	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		return err
	}

	//変数pricesの値を、TaxIncludedPriceJob structsのInputPricesに格納
	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool, errorChan chan error) {
	//インプットデータ（prices.txtの内容）を取得
	err := job.LoadData()

	if err != nil {
		//return err
		//変数errを変数error
		errorChan <- err
		return
	}

	result := make(map[string]string)

	//税込価格を計算し、「税抜き価格：税込価格」を変数resultに格納するfor文
	//job.InputPrices...LoadData()メソッドによって、値がセットされた、TaxIncludedPriceJob structsの値
	for _, price := range job.InputPrices {
		//税込価格を算出
		taxIncludedPrice := price * (1 + job.TaxRate)
		//キー（price）と、それに紐づく値（taxIncludedPrice）を変数resultにセットで格納
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	// WriteJSON...
	// 第一引数...JSONに記載する内容が記載された元ファイル
	// 第二引数...prices.goのTaxIncludedPriceJob structsの全ての変数
	//fmt.Sprintf...指定したフォーマットに従って数値を文字列として出力できる
	job.IOManager.WriteResult(job)
	doneChan <- true
}

// 「税抜き価格」と「税率」を返すメソッド
func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	//変数InputPrices, TaxRateに値を代入
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
