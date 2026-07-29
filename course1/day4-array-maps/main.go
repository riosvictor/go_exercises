package main

import (
	"fmt"
	"slices"
)

func main() {
	// Array
	var array [5]int
	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 4
	array[4] = 5

	// Slice
	slice := []int{1, 2, 3, 4, 5}

	// Map
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	fmt.Println("Array:", array, "Length:", len(array))
	fmt.Println("Slice:", slice, "Length:", len(slice))
	fmt.Println("Map:", m, "Length:", len(m))
	//
	// Slices
	var s []string
	s = make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("Slices >>>")
	fmt.Println("Slice:", s, "Length:", len(s))

	// Append
	s = append(s, "d")
	fmt.Println("Slice:", s, "Length:", len(s))

	// Copy
	s2 := make([]string, len(s))
	copy(s2, s)
	fmt.Println("Slice 2:", s2, "Length:", len(s2))

	// Delete
	s = slices.Delete(s, 0, 1)
	fmt.Println("Slice:", s, "Length:", len(s))

	// Equals
	s1 := []string{"a", "b", "c"}
	s2 = []string{"a", "b", "c"}
	fmt.Println("Slices are equal:", slices.Equal(s1, s2))

	// Maps
	fmt.Println("Maps >>>")

	// Create a map
	m1 := make(map[string]int)
	m1["k1"] = 1
	m1["k2"] = 2
	m1["k3"] = 3

	fmt.Println("Map:", m1, "Length:", len(m1))

	// Delete a key
	delete(m1, "k1")
	fmt.Println("Map:", m1, "Length:", len(m1))

	// Check if a key isKeyPresent
	value, isKeyPresent := m1["k2"]
	if isKeyPresent {
		fmt.Println("Key k2 exists", value)
	} else {
		fmt.Println("Key k2 does not exist")
	}

	//
	sliceInt := []int{1, 2, 3, 4, 5}
	sumSlice := 0
	for _, v := range sliceInt {
		sumSlice += v
	}
	fmt.Println("Sum of slice:", sumSlice)

	//
	mapInt := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("Map a:", mapInt["a"], "Map b:", mapInt["b"])
}
