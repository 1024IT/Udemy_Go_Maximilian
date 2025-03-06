package main

import "fmt"

func pointers() {
	age := 32

	var agePointer *int

	//&...その変数に格納された値を取得できる
	agePointer = &age

	//*...その変数に格納された値のアドレスを取得できる。
	fmt.Println("Age:", *agePointer)

	// adultYears := getAdultYears(age)
	// fmt.Println(adultYears)
}

func getAdultYears(age int) int {
	return age - 18
}
