package main

import "fmt"

func main() {
	x := Soma(1, 2, 3)
	y := multiplica(10, 10)

	fmt.Println(x)
	fmt.Println(y)
}

func Soma(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func multiplica(nums ...int) int {
	total := 1
	for _, n := range nums {
		total *= n
	}
	return total
}
