package main

import "fmt"

func pointers() {
	age := 32

	var agePointer *int

	agePointer = &age

	fmt.Println("Age:", *agePointer)

	// adultYears := getAdultYears(age)
	// fmt.Println(adultYears)
}

func getAdultYears(age int) int {
	return age - 18
}
