package main

import "fmt"

func product(x int, y int, z int) int {
	var product int
	product = x * y * z
	return product
}

func main() {

	//var (
		//product int
		//a, b, c = 3, 6, 3
	//)
	//product = a * b * c
	//fmt.Println(product)

	
	
	fmt.Println(product(2, 4, 6))

}
