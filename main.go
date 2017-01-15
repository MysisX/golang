package main

import (
	"errors"
	"fmt"
)

func isTest(number int) (result bool) {
	defer fmt.Println("done.")
	if number < 0 {
		panic(errors.New("The number is a negative number!"))
	}
	if number%2 == 0 {
		return true
	}
	return
}

func main() {
	fmt.Println(isTest(1))
	fmt.Println("hehe")
}
