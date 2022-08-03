package main

import (
	"fmt"
)

// Slice, map
func main() {
	numbers := []int{1, 2, 3, 4, 5}
	number := numbers[:3]
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
	copy_func(number)
	type_access()
	variadic_func()

	//Calling interface
	var iFe interFace
	iFe = Go{Name: "Vimal"}
	iFe.printName()
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

//Interface is a custom type that is used to specify a set of one or more method signatures and the interface is abstract, so you are not allowed to create an instance of the interface.
type interFace interface {
	printName()
}

// Variadic function is like spread operator in JS
func variadic_func() {
	fmt.Println(calculation("Rectangle", 10, 20))

}

func calculation(shape string, values ...int) int {
	area := 1
	for _, value := range values {
		if shape == "Rectangle" {
			area *= value
		}
	}
	return area

}

func (name Go) printName() {
	fmt.Println("Show name using interface method and struct type : ", name.Name)
}
