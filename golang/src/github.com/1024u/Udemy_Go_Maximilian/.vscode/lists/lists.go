package lists

import (
	"fmt"
)

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[1])
	prices[1] = 9.99

	//append()...スライスに要素を追加する。
	//			　第一引数...スライス
	//			  第二引数...スライスに追加する要素
	updatedPrices := append(prices, 5.99, 12.99, 100.10)
	prices = prices[1:]
	fmt.Println(updatedPrices)

	discountPrices := []float64{101.99, 80.99, 20.59}
	//「discountPrices...」...変数discountPricesの中の全ての値を変数pricesに追加する。
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}

// func main() {
// 	var productNames [4]string = [4]string{"A Book"}
// 	//float64型の4つの値をarrayとして{} 内に入れ、変数pricesに格納
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
// 	fmt.Println(prices)
// 	fmt.Println(productNames)

// 	//変数productNamesの中のArrayの2つ目のデータを"A Carpet"に変更
// 	productNames[2] = "A Carpet"

// 	//arrayが格納された変数pricesの中の2つ目のデータを出力
// 	fmt.Println(prices[2])

// 	//変数pricesに格納された、1つ目以上〜3つ目未満の値を、変数featuredPricesに格納(Slices)
// 	//※補足：[1:]や[:3]のように、「下限値以上の全ての値」「上限値未満の全ての値」という書き方も可能
// 	featuredPrices := prices[1:3]
// 	fmt.Println(featuredPrices)

// 	//len()...格納されているデータの数を返す
// 	//cap()...スライスが持つことのできる最大の要素数を返す（スライスの宣言時に最大の要素数を定義できる）
// 	fmt.Println(len(featuredPrices), cap(featuredPrices))
// }
