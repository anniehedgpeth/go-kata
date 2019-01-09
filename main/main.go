package main

import "fmt"

var name = "Annie"

func isOwner(name string) bool {
	return name == "Annie"
}

func sayHi(name string) {
	fmt.Printf("Hello, %v!\n\n", name)
	if isOwner(name) {
		fmt.Println("Your name is like the owner of this repo!")
	}
}

func main() {
	sayHi(name)
	familyMembers := []string{"Annie", "Michael", "Kid1", "Kid2", "Kid3"}
	for i := 0; i < len(familyMembers); i++ {
		item := familyMembers[i]
		fmt.Printf("%d: %v\n", i, item)
	}
}
