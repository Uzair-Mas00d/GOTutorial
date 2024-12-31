package main

import "fmt"

func sayHello(s string) {
	s = "Hello World"
}

func sayHelloPointer(s *string) {
	*s = "Hello World"
}

func main() {
	var greeting string = "Hello Go"
	sayHello(greeting)
	fmt.Println("After Calling sayHello", greeting)

	sayHelloPointer(&greeting)
	fmt.Println("After Calling sayHelloPointer", greeting)

	var intPointer *int
	fmt.Println(intPointer)

	isFast := false
	isSlow := true
	if isFast {
		fmt.Println("Your are going fast")
	}else if isSlow {
		fmt.Println("Your going slow")
	}else {
		fmt.Println("Thanks")
	}

	const upperSpeedLimit, lowerSpeedLimit int = 120, 80
	speed := 120
	if speed >= upperSpeedLimit {
		fmt.Println("Your are going fast")
	} else if speed <= lowerSpeedLimit {
		fmt.Println("Your are going slow")
	}else {
		fmt.Println("Thanks")
	}

	isWeekend := false
	isHoliday := false
	if isWeekend && isHoliday {
		fmt.Println("Hoilday on weekend")
	}else if isWeekend && !isHoliday {
		fmt.Println("wwekend")
	}else if !isWeekend && isHoliday {
		fmt.Println("Holiday")
	} else {
		fmt.Println("weekday")
	}

	var weekend bool =true
	var hoilday bool = false
	isDayOff := weekend || hoilday
	if isDayOff {
		fmt.Println("dayoff")
	}
	isHoildayWeekend := weekend && hoilday
	if isHoildayWeekend {
		fmt.Println("holiday weekend")
	}
}
