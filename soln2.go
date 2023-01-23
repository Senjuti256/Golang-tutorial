package main

import "fmt"

func main() {

	const (
		Nov = 11
		Oct = 10
		Sep = 9
	)
	fmt.Println(Nov, Oct, Sep)
	fmt.Println("--------------------")
	const (
		Nov1 = (11 - iota)
		Oct1
		Sep1
	)
	fmt.Println(Nov1, Oct1, Sep1)
	fmt.Println("--------------------")

}
