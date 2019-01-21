package main

import "fmt"

var name = "Annie"

func main() {
	sayHi()
	isOwner()
	theSimpsons := []string{"Homer", "Marge", "Bart", "Lisa", "Maggie"}
	// It's easier to use "i" for index instead of something like "name" - less confusing.
	for i := 0; i < len(theSimpsons); i++ {
		person := theSimpsons[i]
		fmt.Printf("%d: %v\n", i, person)
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
