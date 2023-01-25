package main

import (
	"fmt"
	"GreetingsApp/pack"
)

var s string

func main(){
	
	fmt.Println("Welcome to my app!")
	fmt.Println("........................")
	fmt.Println("Enter your name:")
	fmt.Scanln(&s)
	pack.Hello(s)
}
