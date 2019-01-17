package main

import "fmt"
var name = "Annie"

func main() {
	sayHi()
	isOwner()
	theSimpsons := []string{"Homer", "Marge", "Bart", "Lisa", "Maggie"}
	for i, item := range theSimpsons {
		fmt.Printf("%v: %v\n", i, item)
	}
}

func sayHi() {
	fmt.Printf("Hello, %v!\n", name)

}

func isOwner() {
	if name == "Annie" {
		fmt.Println("Your name is like the owner of this repo!")
	}
}