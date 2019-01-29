package main

import "fmt"

var name = "Annie"
var members = []string{"Homer", "Marge", "Bart", "Lisa", "Maggie"}
var familyMemberMap = map[string]string{"Homer": "blue", "Marge": "purple", "Bart": "black", "Lisa": "red", "Maggie": "white"}

func main() {
	sayHi()
	isOwner()
	fmt.Println()
	listMembersIndex()
	fmt.Println()
	listMembersRange()
	fmt.Println()
	sayFavoriteColors()
}

func sayHi() {
	fmt.Printf("Hello, %v!\n", name)
}

func isOwner() {
	if name == "Annie" {
		fmt.Println("Your name is like the owner of this repo!")
	}
}

func listMembersIndex() {
	for i := 0; i < len(members); i++ {
		item := members[i]
		fmt.Printf("%d: %v\n", i, item )
	}
}

func listMembersRange() {
	for i, item := range members {
		fmt.Printf("%d: %v\n", i, item )
	}
}

func sayFavoriteColors() {
	for name, color := range familyMemberMap {
		fmt.Printf("%v: %v\n", name, color)
	}
}