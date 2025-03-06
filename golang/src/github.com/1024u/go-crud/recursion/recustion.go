//学べること...
// Recursion

package recursion

import "fmt"

func main() {
	fact := factorial(5)
	fmt.Println(fact)
}

func factorial(number int) int {
	//Recursionを使った5の階乗
	//factorial(5) から factorial(0) まで「呼び出し」が進む（順方向）。
	//factorial(0) = 1 になったら、「結果を返しながら」戻っていく（逆方向）。
	//最終的に factorial(5) = 120 が確定して main() に返される。

	if number == 0 {
		return 1
	}

	return number * factorial(number-1)

	//Recursionを使わない5の階乗
	// result := 1

	// for i := 1; i <= number; i++ {
	// 	result = result * i
	// }

	// return result
}

// 5の階乗...5 * 4 * 3 * 2 * 1 = 120
//※階乗...1から与えられた整数までのすべての数の積。
