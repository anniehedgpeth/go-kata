package main

import "fmt"

var name = "Annie"

func sayHi(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func isOwner(name string) bool {
	if name == "Annie" {
		fmt.Println("Your name is is like the owner of this repo!")
	}
	return name == "Annie"
}

func main() {
	sayHi(name)
	familyMembers := map[string]string{"Homer": "blue", "Marge": "purple", "Bart": "black", "Lisa": "red", "Maggie": "white"}
	for name, color := range familyMembers {
		fmt.Printf("%v: %v\n", name, color)
	}
	isOwner(name)
}
