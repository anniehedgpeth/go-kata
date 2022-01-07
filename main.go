package main

import (
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

	simpsons := []person{
		person {
			"Marge",
			"red",
			50.50,
			1,
		},
		person {
			"Homer",
			"green",
			50.25,
			2,
		},
		person {
			"Bart",
			"purple",
			20.20,
			12,
		},
		person {
			"Lisa",
			"blue",
			20.25,
			25,
		},
		person {
			"Maggie",
			"yellow",
			3.25,
			3,
		},
	}

	for _, p := range simpsons {
		greeting.SayHi(p.name)
		greeting.SayBirthdays(p.birthday)
		greeting.SayAllowances(p.allowance)
		greeting.FavColor(p.color)
	}
}