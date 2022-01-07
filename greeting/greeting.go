package greeting

import (
	"fmt"
)

func PrintDivider() {
	fmt.Println("\n**************************************************")
}

func SayHi(name string) {
	fmt.Printf("\nHello, %s!\n", name)
	IsOwner(name)
	PrintDivider()
}

func IsOwner(name string) bool {
	if name == "Annie" {
		fmt.Println("\nYour name is like the owner of this repo!")
		return true
	}
	return false
}

func SayAllowances(allowances map[string]float32) {
	for m, a := range allowances {
		fmt.Printf("%v: %.02f\n", m, a)
	}
	PrintDivider()
}

func SayBirthdays(birthdays map[string]int) {
	for d, b := range birthdays {
		fmt.Printf("%v: %d\n", d, b)
	}
	PrintDivider()
}

func FavColor(simpsons map[string]string) {
	for n, s := range simpsons {
		fmt.Printf("%v: %v\n", n, s)
	}
	PrintDivider()
}
