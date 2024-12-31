package main

import (
	"fmt"
	"strings"
)

func main() {
	color := "blue"

	switch color {
	case "red":
		fmt.Println("Red")
	case "blue":
		fmt.Println("Blue")
	default:
		fmt.Println("Unknown")
	}

	day := "Tue"
	switch day {
	case "Mon", "Tue", "Wed", "Thu", "Fri":
		fmt.Println("Weekday")
	case "Sat", "Sun":
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid Day")
	}

	age := 25
	switch {
	case age < 18:
		fmt.Println("Minor")
	case age > 18 && age < 60:
		fmt.Println("Adult")
	case age > 60:
		fmt.Println("Old")
	}

	useRole := "admin"
	switch useRole {
	case "admin":
		fmt.Println("Full Access")
		fallthrough
	case "editor":
		fmt.Println("Publish Content")
		fallthrough
	case "contributor":
		fmt.Println("Write Content")
		fallthrough
	case "viewr":
		fmt.Println("View Content")
	default:
		fmt.Println("No Access")
	}

	var temp float64
	var fromUnit, toUnit string

	fmt.Println("Enter Temprature: ")
	fmt.Scanln(&temp)

	fmt.Println("Enter current unit (C, K, F): ")
	fmt.Scanln(&fromUnit)
	fromUnit = strings.ToUpper(fromUnit)

	fmt.Println("Enter unit to convert to (C, K, F): ")
	fmt.Scanln(&toUnit)
	toUnit = strings.ToUpper(toUnit)

	convertedTemp := 0.0

	switch fromUnit {
	case "C":
		if toUnit == "F" {
			convertedTemp = temp * 9/5 + 32
		} else if toUnit == "K" {
			convertedTemp = temp + 273.15
		}
	case "F":
		if toUnit == "C" {
			convertedTemp = (temp - 32) * 5/9
		} else if toUnit == "K" {
			convertedTemp = (temp - 32) * 5/9 + 273.15
		}
	case "K":
		if toUnit == "C" {
			convertedTemp = temp - 273.15
		} else if toUnit == "F" {
			convertedTemp = (temp - 273.15) * 9/5 + 32
		}

	}

	fmt.Printf("Converted Temprature %.2f %s\n",convertedTemp, toUnit)

}
