package greeting

import "fmt"

var name = "Annie"

func SayHi() {
	fmt.Printf("Hello, %v!\n", name)
	if isOwner(name) {
		fmt.Println("Your name is like the owner of this repo!")
	}
}

func isOwner(name string) bool {
	return name == "Annie"
}

func SayFavoriteColorsOfSimpsons() {
	for _, person := range simpsons {
		fmt.Printf("%v: %v\n", person.name, person.color)
	}
}

func SayBirthdaysOfSimpsons() {
	for _, person := range simpsons {
		fmt.Printf("%v: %d\n", person.name, person.birthdate)
	}
}

func SayAllowanceOfSimpsons() {
	for _, person := range simpsons {
		fmt.Printf("%v: %2.2f\n", person.name, person.allowance)
	}
}

type person struct {
	name      string
	birthdate int
	allowance float32
	color     string
}

var simpsons = []person{
	{name: "Homer", birthdate: 22, allowance: 10.50, color: "blue"},
	{name: "Marge", birthdate: 14, allowance: 20.34, color: "red"},
	{name: "Bart", birthdate: 1, allowance: 5.50, color: "black"},
	{name: "Lisa", birthdate: 4, allowance: 15.50, color: "purple"},
	{name: "Maggie", birthdate: 1, allowance: 6.50, color: "white"},
}
