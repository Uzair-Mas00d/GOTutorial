package main

import "fmt"

func main() {
	carInventory := make(map[string]int)
	carInventory["Sedan"] = 25
	carInventory["SUV"] = 15
	carInventory["Convertible"] = 10

	fmt.Println("Car Inventory", carInventory)

	carInv := map[string]int {
		"Sedan": 25,
		"SUV": 15,
		"Convertible": 10,
	} 
	fmt.Println(carInv)

	fmt.Println(len(carInv))

	fmt.Println(carInv["Sedan"])
	fmt.Println(carInv["Cope"])

	carInv["SUV"] = 50
	fmt.Println(carInv)

	delete(carInv, "Convertible")

	fmt.Println(carInv)

	numConvertibles, convertibleFound := carInv["Convertible"]
	fmt.Println("Convertible Found", convertibleFound)
	if convertibleFound {
		fmt.Println("Convertible", numConvertibles)
	}

	if numConvertible, ok := carInv["Convertible"]; ok {
		fmt.Println("Convertible", numConvertible)
	}

	clear(carInv)
	fmt.Println(carInv)

	bodyTypes := [3]string{"Sedan", "SUV", "Convertible"}

	for i, bodyType := range bodyTypes {
		fmt.Printf("Index %v Item %v\n", i, bodyType)
	}

	for i := range bodyTypes {
		fmt.Printf("Index %v\n", i)
	}

	for _, bodyType := range bodyTypes {
		fmt.Printf("Item %v\n", bodyType)
	}

	for bodyType, count := range carInventory {
		fmt.Printf("%v -> %v\n", bodyType, count)
	}

	for bodyType:= range carInventory {
		fmt.Printf("%v\n", bodyType)
	}

	for _, count := range carInventory {
		fmt.Printf("%v\n", count)
	}

}
