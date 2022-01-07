package greeting

import (
	"fmt"
	"github.com/anniehedgpeth/go-kata/greeting"
)

type person struct {
	name string
	color string
	allowance float32
	birthday int
}

func main() {
	name := "Annie"
	greeting.SayHi(name)

	simpsons := map[string]string{
		"Marge":  "red",
		"Homer":  "blue",
		"Bart":   "yellow",
		"Lisa":   "purple",
		"Maggie": "orange",
	}

	greeting.FavColor(simpsons)

	birthdays := map[string]int{
		"Marge":  1,
		"Homer":  2,
		"Bart":   12,
		"Lisa":   25,
		"Maggie": 3,
	}

	greeting.SayBirthdays(birthdays)

	allowances := map[string]float32{
		"Marge":  50.50,
		"Homer":  50.25,
		"Bart":   12.32,
		"Lisa":   12.33,
		"Maggie": 3,
	}

	greeting.SayAllowances(allowances)
}