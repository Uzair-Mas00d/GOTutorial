package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println("Hello World")

	var agency string = "Fast Track"
	fmt.Println("Welcome to", agency)

	name := "Ali"
	fmt.Println("Hello", name)

	var totalCars int = 50
	fmt.Println(totalCars, "Cars")

	startingPrice := 29.99
	fmt.Printf("Our prices start at %v\n", startingPrice)
	fmt.Println("Take your pick")

	// var insuranceInclude bool = true
	insuranceInclude := true
	println("Insurance Include", insuranceInclude)

	var names string
	fmt.Println("name", names)
	var count int
	fmt.Println("count", count)
	var price float64
	fmt.Printf("price %.1f\n", price)
	var insured bool
	fmt.Println("insured", insured)

	str := "Hello World"
	length := len(str)
	fmt.Println(length)

	str1 := "Go Program"
	str2 := "go Program"
	fmt.Println(strings.EqualFold(str1, str2))

	str3 := "Hello Wolrd"
	wIndex := strings.Index(str3, "W")
	fmt.Println(wIndex)
	fmt.Println(strings.Index(str3, "Hello"))
	subStr := str3[wIndex: 11]
	fmt.Println(subStr)

	str4 := "Hello Go"
	fmt.Println(strings.Replace(str4, "Go", "World", 1)) 

	str5 := "Go Pogram"
	fmt.Println(strings.ToUpper(str5)) 
	fmt.Println(strings.ToLower(str5)) 
	fmt.Println(strings.Contains(str5, "Go"))

	tempC := 32.0
	tempK := 0.0

	tempK = tempC + 273.15
	tempF := tempC * 9/5 + 32

	fmt.Println("C",tempC)
	fmt.Println("K",tempK)
	fmt.Println("F", tempF)

	var tempratureC float64 = 12.75
	fmt.Println(math.Round(tempratureC))
	fmt.Println(math.Ceil(tempratureC))
	fmt.Println(math.Floor(tempratureC))
	fmt.Println(math.Abs(-5.5))
	fmt.Println(math.Pow(3,2))
	fmt.Println(math.Sqrt(16))
	fmt.Println(math.Sin(math.Pi/2))
	fmt.Println(math.Cos(math.Pi))
	fmt.Println(math.Exp(1))
	fmt.Println(math.Log(math.E))
}
