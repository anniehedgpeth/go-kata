package greeting

import (
	"fmt"
)

func PrintDivider() {
	fmt.Println("\n**************************************************")
}

func SayHi(name string) {
	PrintDivider()
	fmt.Printf("\nHello, %s!\n", name)
	IsOwner(name)
}

func IsOwner(name string) bool {
	if name == "Annie" {
		fmt.Println("\nYour name is like the owner of this repo!")
		return true
	}
	return false
}

func SayAllowances(allowance float32) {
	fmt.Printf("allowance: %.02f\n", allowance)
}

func SayBirthdays(birthday int) {
	fmt.Printf("birthday: %d\n", birthday)
}

func FavColor(color string) {
	fmt.Printf("color: %v\n", color)
	PrintDivider()
}
