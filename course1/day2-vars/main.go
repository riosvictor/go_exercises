package main

import "fmt"

func main() {
	// Declare a variable
	var name string
	// Assign a value to the variable
	name = "John Doe"
	var age int = 30
	// Infer the type of the variable
	city := "São Paulo"
	// Multiple variable declaration
	var country, zipCode = "Brazil", "12345-678"
	// Block variable declaration
	var (
		// Declare a variable
		state string = "SP"
		// Assign a value to the variable
		neighborhood = "Jardim Paulista"
	)
	// Var without value
	var number int
	// Constant variable declaration
	const pi = 3.14
	// Shadow variable declaration
	var anotherName = "Jane Doe"
	// Literal variable declaration
	var height = 1.75

	// Print the variable
	fmt.Println("Olá", name)
	fmt.Println("Idade:", age)
	fmt.Println("Cidade:", city)
	fmt.Println("País:", country)
	fmt.Println("CEP:", zipCode)
	fmt.Println("Estado:", state)
	fmt.Println("Bairro:", neighborhood)
	fmt.Println("Número:", number)
	fmt.Println("Pi:", pi)
	fmt.Println("Outro nome:", anotherName)
	// Shadow variable
	anotherName = "John Smith"
	fmt.Println("Outro nome:", anotherName)
	fmt.Println("Altura:", height)
}
