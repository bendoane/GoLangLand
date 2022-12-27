package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello World, let me introduce...")
	friend := Person{"Bob", "TheBuilder", 42}
	pupper := Dog{"Casius the grand", "English Mastif", 4}

	fmt.Println(introduceHuman(friend))
	fmt.Println("And")
	fmt.Println(introduceDoggo(pupper))
}

type Person struct {
	firstName string
	lastName  string
	age       int
}

type Dog struct {
	name  string
	breed string
	age   int
}

func introduceHuman(p Person) string {
	return p.firstName + " " + p.lastName + ", who is " + strconv.Itoa(p.age) + " years old."
}

func introduceDoggo(d Dog) string {
	return d.name + ", the " + d.breed + ", who is " + strconv.Itoa(d.age) + " years old."
}
