package main

import (
	"fmt"
)

var name = "Levi"

func main(){
	sayHi()
}

func sayHi() {
	fmt.Printf("Hello, %v!\n", name)
	if name == "Levi" {
		fmt.Println("Your name is like the owner of this repo!")
	}

}
