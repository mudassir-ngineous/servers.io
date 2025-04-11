package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

type Person struct {
	Name string
	Age int
}

// implement the Speaker interface for Dog and Person

func (d Dog) Speak() string {
	return "Woof!"
}

func (p Person) Speak() string {
	return "Hello, my name is " + p.Name
}

func makeItSpeak(s Speaker) {
    fmt.Println(s.Speak())
}

func main() {
	p := Person{Name: "Mudassir"}
    d := Dog{Name: "Doggy"}

    makeItSpeak(p) // Output: Hello, I'm Mudassir
    makeItSpeak(d) // Output: Woof, I'm Doggy
}




