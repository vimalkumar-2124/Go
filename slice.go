package main

import (
	"fmt"
)

// Slice, map
func main() {
	numbers := []int{1, 2, 3, 4, 5}
	number := numbers[:3]
	copy_func(number)
	type_access()
	dict := make(map[string]string)
	dict["name"] = "Vimal"
	dict["age"] = "24"
	fmt.Println(dict["name"])
	//For loop used with range
	for _, i := range numbers {
		fmt.Println("With range : ", i)
	}
	//Normal for loop
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Without range : ", numbers[i])
	}
	i := 0
	//Infinite loop
	for {
		fmt.Println("Infinite loop : ", numbers[i])
		if i == 3 {
			break
		}
		i++
	}
}

//Copy
func copy_func(number []int) {
	numberCopy := make([]int, len(number))
	copy(numberCopy, number)
	fmt.Printf("Numbers %d \n", number)
	fmt.Printf("Number %d \n", numberCopy)
	result := append_func(number, numberCopy)
	fmt.Printf("Appended arr : %T %d \n", result, result)
}

//Append
func append_func(number []int, numberCopy []int) []int {
	arr := append(number, numberCopy...)
	return arr

}

//struct
type Go struct {
	Name string
	age  string
}

func type_access() {
	var obj Go
	obj.Name = "Vimal"
	obj.age = "24"
	fmt.Printf("Name : %v \n", obj.Name)
}
