package main

import (
	"errors"
	"fmt"
)

// example: struct
type Person struct {
    Name string
    Age  int
}

func main() {
	fmt.Println("--------------------------------")
	fmt.Println("Hello, World!")

	// Basic types: int, float64, string, bool
	// Zero values: Variables declared without initialization get a default (e.g., 0 for int, "" for string).
	var x int = 10        // Explicit declaration
    y := 20               // Short declaration (type inferred)
    name := "abc-xyz"     // String
    isCool := true        // Boolean
    pi := 3.14            // Float

    fmt.Println("Basic types example: x, y, name, isCool, pi:\t", x, y, name, isCool, pi)

	fmt.Println("--------------------------------")
    if x > 5 {		// no parentheses required for if statement
        fmt.Println("If statement example: x is big")
    } else {
        fmt.Println("If statement example: x is small")
    }

	// Only for is used for loops (no while).
	fmt.Println("--------------------------------")
    for i := 0; i < 3; i++ {	// no parentheses required for for loop
        fmt.Println("For loop example: i:\t", i)
    }

	// example: function call
	fmt.Println("--------------------------------")
	result := add(3, 4)
    fmt.Println("Function call example: result:\t", result)

	// example: multiple return values
	fmt.Println("--------------------------------")
	fmt.Println("Hello", "World")
	s1, s2 := swap("Hello", "World")
    fmt.Println("Multiple return values example: s1, s2:\t", s1, s2)

	// example: array
	fmt.Println("--------------------------------")
	var arr1 [3]int			// method 1: declare an array
	var arr2 [3]string = [3]string{"foo", "bar", "baz"}	// method 2: declare and initialize an array
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3

	fmt.Println("Array example: arr1:\t",arr1)
	fmt.Println("Array example: arr2:\t",arr2)

	fmt.Println("--------------------------------")
	slice := []int{1, 2, 3}		// declare and initialize
	slice = append(slice, 4) 	// append value 4 to slice array and assign back to slice
	fmt.Println("slice:\t",slice)

	slice2 := slice[1:3] // Slicing: [2 3]
	fmt.Println("Slice example: slice2:\t", slice2)

	// example: map
	// The syntax is a short declaration of a map in Go. It creates a map m with keys of type string and values of type int.
	fmt.Println("--------------------------------")
	m := map[string]int{"age": 30, "score": 100}
	fmt.Println("Map example: m[\"age\"]:\t", m["age"]) // 30

	// example: struct
	// Go uses structs to create custom types
	fmt.Println("--------------------------------")
	p := Person{Name: "John", Age: 30}
	fmt.Println("Struct example: p:\t", p.Name, p.Age)

	
	// example: pointers
	// Pointers are used to store the memory address of a variable.
	// The syntax is a short declaration of a pointer in Go. It creates a pointer p to a Person struct.
	fmt.Println("--------------------------------")
	xyz := 10
	pqr := &xyz       // pqr is a pointer to xyz
	*pqr = 20       	// Changes xyz to 20
	fmt.Println("Pointer example: xyz:\t", xyz) // 20

	fmt.Println("--------------------------------")
	result_divide, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error example: err:\t", err)
	} else {
		fmt.Println("Result example: result_divide:\t", result_divide)
	}
	//
}

// add adds two integers and returns their sum.
// It takes two parameters:
// - a: the first integer
// - b: the second integer
// Returns the sum of a and b.
func add(a int, b int) int {
    return a + b
}

// swap swaps two strings and returns them.
// It takes two parameters:
// - x: the first string
// - y: the second string
// Returns the two strings swapped.
func swap(x, y string) (string, string) {
    return y, x
}

func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("divide by zero")
    }
    return a / b, nil
}