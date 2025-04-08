package main

import "fmt"

func main() {
	instance = GetInstance()
	calculatorInstance := GetInstance()

	if instance != calculatorInstance {
		panic("Instance is not the same")
	}

	fmt.Println("Instance is the same")
}