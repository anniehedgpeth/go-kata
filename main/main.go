package main

import "fmt"

var name = "Annie"

func main() {
	sayHi()
	isOwner()
	simpsonsMember := []string{"Homer", "Marge", "Bart", "Lisa", "Maggie"}
	for name := 0; name < len(simpsonsMember); name++ {
		person := simpsonsMember[name]
		fmt.Printf("%d: %v\n", name, person)
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
