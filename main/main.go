package main

import "fmt"

var name = "Annie"
var simpsons = []string{"Homer", "Marge", "Bart", "Lisa", "Maggie"}
var simpsonsColors = map[string]string{
	"Homer":  "blue",
	"Marge":  "green",
	"Bart":   "black",
	"Lisa":   "red",
	"Maggie": "yellow"}

var birthdays = map[string]int{
	"Homer":  1,
	"Marge":  2,
	"Bart":   3,
	"Lisa":   4,
	"Maggie": 5}

var allowances = map[string]float64{
	"Homer":  1.25,
	"Marge":  2.25,
	"Bart":   3.50,
	"Lisa":   4.50,
	"Maggie": 5.75}

func main() {
	sayHi("Annie")
	isOwner()
	fmt.Println()
	listSimpsons()
	fmt.Println()
	listSimpsonsRangeWithIndex()
	fmt.Println()
	listSimpsonsRange()
	fmt.Println()
	sayFavoriteColors()
	fmt.Println()
	sayBirthdays()
	fmt.Println()
	sayAllowances()
}

func sayHi(s string) {
	fmt.Printf("Hello, %v!\n", s)
}

func isOwner() {
	if name == "Annie" {
		fmt.Println("Your name is the same as the owner of this repo!")
	}
}

func listSimpsons() {
	fmt.Println("*** CLASSIC FOR LOOP ***")
	for i := 0; i < len(simpsons); i++ {
		item := simpsons[i]
		fmt.Printf("%d: %v\n", i, item)
	}
}

func listSimpsonsRangeWithIndex() {
	fmt.Println("*** RANGE WITH INDEX ***")
	for i, item := range simpsons {
		fmt.Printf("%d: %v\n", i, item)
	}
}

func listSimpsonsRange() {
	fmt.Println("*** RANGE ***")
	for _, item := range simpsons {
		fmt.Printf("%v\n", item)
	}
}

func sayFavoriteColors() {
	fmt.Println("*** COLOR RANGE ***")
	for k, v := range simpsonsColors {
		fmt.Printf("%s: %v\n", k, v)
	}
}

func sayBirthdays() {
	fmt.Println("*** INTEGER RANGE ***")
	for k, v := range birthdays {
		fmt.Printf("%s: %v\n", k, v)
	}
}

func sayAllowances() {
	fmt.Println("*** FLOAT RANGE ***")
	for k, v := range allowances {
		fmt.Printf("%s: $%v\n", k, v)
	}
}
