package maps

import "fmt"

func main() {
	//map[string]string...map[キーのデータ型]キーに紐づく値のデータ型{}
	//structsだと、値を足したり削除したりできないが、mapはできる。
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)

	//"Amazon Web Services"というキーに紐づいた値を出力
	fmt.Println(websites["Amazon Web Services"])

	//新しいキーと、それに紐づく値を変数websitesに追加
	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)

	//変数websitesから、キー"Google"と、それに紐づく値を削除
	delete(websites, "Google")
	fmt.Println(websites)

	// websites := []string{"https://google.com", "https://aws.com"}
}
