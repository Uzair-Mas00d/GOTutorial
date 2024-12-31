package main

import "fmt"

func main() {
	var bodyTypes [3]string

	bodyTypes[0] = "Sedan"
	bodyTypes[1] = "SUV"
	bodyTypes[2] = "Convertible"

	fmt.Println(bodyTypes)

	bodyType := [3]string{"Sedan", "SUV", "Convertible"}

	fmt.Println(bodyType[0])
	fmt.Println(bodyType[1])
	fmt.Println(bodyType[2])

	var carFleet [3][2]string
	carFleet[0] = [2]string{"2 Sedan Avaible", "2 Sedan Booked"}
	carFleet[1] = [2]string{"3 SUV Avaible", "4 SUV Booked"}
	carFleet[2] = [2]string{"1 Convertible Avaible", "1 Convertible Booked"}

	fmt.Println(carFleet)

	for i := 0; i < len(carFleet); i++ {
		row := carFleet[i]
		for j:=0; j < len(row); j++ {
			fmt.Printf("%v ",row[j])
		}
		fmt.Println()
	}

	var fuelTypes []string
	fmt.Println(len(fuelTypes))

	fmt.Println("The slice is nil:", fuelTypes == nil)

	fuelType := []string{"Electric", "Gasoline", "Hybrid"}
	fmt.Println(fuelType)

	fuelType = append(fuelType, "Diesel", "Hydrogen")
	fmt.Println(fuelType)

	fuelTYPES := make([]string, 3)
	fmt.Println(len(fuelTYPES))
	fmt.Println("The slice is nil:", fuelTYPES == nil)
	fuelTYPES[0] = "Eletric"
	fuelTYPES[1] = "Gasoline"

	fuelTYPES = append(fuelTYPES, "Hybrid", "Diesel", "Hydrogen")

	fmt.Println(fuelTYPES)

	fuelTypeCopy := make([]string, len(fuelTYPES))
	copy(fuelTypeCopy, fuelTYPES)

	fmt.Println(fuelTypeCopy)

	popular := fuelTYPES[0:2]
	fmt.Println("popular", popular) 

	clean := fuelTYPES[3:]
	fmt.Println("popular", clean) 
}
