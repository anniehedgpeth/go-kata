package main

import "fmt"

var name = "Annie"

func sayHi() {
	fmt.Printf("Hello, %v!\n", name)
}

func isOwner() {
	if name == "Annie" {
		fmt.Println("Your name is like the owner repo!")
	}
}

func sayFavoriteColorsOfSimpsons() {
	familymembers := map[string]string{"Homer": "blue", "Marge": "yellow", "Bart": "black", "Lisa": "red", "Maggie": "purple"}
	for name, color := range familymembers {
		fmt.Printf("%v: %v\n", name, color)
	}
}

func sayBirthdayOfSimpsons() {
	familymembers := map[string]int{"Homer": 1, "Marge": 2, "Bart": 3, "Lisa": 4, "Maggie": 5}
	for name, date := range familymembers {
		fmt.Printf("%v: %d\n", name, date)
	}
}

func sayAllowancesOfSimpsons() {
	familymembers := map[string]float64{"Homer": 10.5, "Marge": 20.6, "Bart": 3.4, "Lisa": 4.5, "Maggie": 2.5}
	for name, allowance := range familymembers {
		fmt.Printf("%v: $%2.2f\n", name, allowance)
	}
}

func main() {
	sayHi()
	sayFavoriteColorsOfSimpsons()
	sayBirthdayOfSimpsons()
	sayAllowancesOfSimpsons()
	isOwner()
}
