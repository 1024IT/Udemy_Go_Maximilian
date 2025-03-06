//学べること...
// 無名関数(anonymous function)...変数に代入された関数。fmt.Println(変数名)で、代入された関数を実行できる。

package anonymous

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	//transformed...無名関数(anonymous function)。
	//				変数に代入された関数。fmt.Println(変数名)で、代入された関数を実行できる。
	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	fmt.Println(transformed)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
