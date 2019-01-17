package main

import "fmt"

func main() {
	sayHi()
}

func sayHi() {
	name := "World"
	fmt.Printf("Hello, %v!", name)
}
