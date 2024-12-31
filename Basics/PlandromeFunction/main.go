package main

import (
	"fmt"
	"math"
	"unicode"
)

var carInvntory = map[string]int{} 

func addInventory (bodyType string, count int){
	carInvntory[bodyType] += count
	fmt.Println(bodyType, "added New Count:", carInvntory[bodyType])
}

func getCount(bodyType string) int {
	fmt.Printf("Looking for body type %v\n", bodyType)
	count := carInvntory[bodyType]

	return count
}

func getStatus () (string, string) {

	return "Hello", "World"
}

func sum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}

	return total
}

func processNumber(num int, opreation func(int) int) int{
	return opreation(num)
}

func double(num int) int{
	return num * 2
}

func main() {
	input := "Madam, I'm Adam"

	normalized := ""
	for _, r := range input {
		if unicode.IsLetter(r){
			normalized += string(unicode.ToLower(r))
		}

	}

	var isPlandrome bool = true

	length := len(normalized)
	for i := 0; i < length/2; i++ {
		if normalized[i] != normalized[length-1-i]{
			isPlandrome = false
			break
		}
	}

	if isPlandrome {
		fmt.Println("Plandrome")
	}else {
		fmt.Println("Not Plandrome")
	}


	addInventory("Sedan", 10)
	println("Found:", getCount("Sedan"))

	hello, world := getStatus()
	fmt.Println(hello, world)

	fmt.Println("Sum:", sum(1,2,3,4,5))

	result := processNumber(3, double)
	fmt.Println(result)

	results := processNumber(5, func (i int) int  {
		return int(math.Pow(float64(i), 2))
	})
	fmt.Println(results)
}