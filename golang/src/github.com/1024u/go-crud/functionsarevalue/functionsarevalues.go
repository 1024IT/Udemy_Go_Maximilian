// 学べること...
//高階関数(Higher-Order Function)...関数を引数として受け取る、または関数を返す関数である。

package main

import (
	"fmt"
)

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}

	moreNumbers := []int{5, 1, 2}
	//double...double()メソッドを引数としてtransformNumbersに渡している
	//		　　引数として渡しているだけなので、まだ実行はしていない。
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	transformerFn1 := getTransformerFunction(&numbers)
	transformerFn2 := getTransformerFunction(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, transformerFn1)
	moreTransformedNumbers := transformNumbers(&moreNumbers, transformerFn2)

	fmt.Println(transformedNumbers)
	fmt.Println(moreTransformedNumbers)
}

// transform func(int) int...高階関数(Higher-Order Function)の1つ。
// 　　　　　　　　　　　　　　　　関数を引数として受け取る、または関数を返す関数である。
func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val)) //ここでdouble()メソッドが実行
	}

	return dNumbers
}

func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
