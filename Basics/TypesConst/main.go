package main

import (
	"fmt"
	"strconv"
)

func main() {
	var tempInt int = 10
	var tempFloat float64 = float64(tempInt)
	fmt.Println(tempFloat)
	fmt.Printf("%.1f\n",tempFloat)

	fmt.Printf("%T\n", tempInt)

	str := fmt.Sprint(80)
	fmt.Println(str)
	fmt.Printf("%T\n", str)

	str2 := strconv.Itoa(800)
	fmt.Println(str2)
	fmt.Printf("%T\n", str2)

	var myStr string = "24"
	var intStr, _ = strconv.Atoi(myStr)
	fmt.Println(intStr)
	fmt.Printf("%T\n", intStr)

	var myFloatStr string = "3.1415"
	var floatStr, _ = strconv.ParseFloat(myFloatStr, 64)
	fmt.Println(floatStr)
	fmt.Printf("%T\n", floatStr)

	mybool, _ := strconv.ParseBool("t")
	fmt.Println(mybool)

	// const Agency string = "Hello World"
	const Agency = "Hello World"
	fmt.Println(Agency)

	const (
		Founded = 2001
		Founder = "Ali"
	)
	fmt.Println(Founded)
	fmt.Println(Founder)

	const (
		_ = iota
		Economy
		Compact
		Standard
	)
	fmt.Println("Economy",Economy)
	fmt.Println("Compact",Compact)
	fmt.Println("Standard",Standard)

	fmt.Println("What is your name")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("Hello %v\n", name)
}
