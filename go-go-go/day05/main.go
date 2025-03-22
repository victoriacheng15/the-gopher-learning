package main

// Day 05: Data Types - floating-point numbers with float32 and float64

import "fmt"

// basic data type - float
// float32 - 32 bits range from -3.4e38 to 3.4e38
// float64 - 64 bits range from -1.7e308 to 1.7e308
// float64 is the default data type for floating-point numbers

func allFloat32() {
	// float32
	var x float32 = 500.75
	var y float32 = 3.4e38

	fmt.Println("Float - float32")
	fmt.Println("Value: ", x) // Type: float32, Value: 500.75
	fmt.Println("Value: ", y) // Type: float32, Value: 3.4e38
}

func allFloat64() {
	// float64
	var a float64 = 500.75
	var b float64 = 1.7e308

	fmt.Println("\nFloat - float64")
	fmt.Println("Value: ", a) // Type: float64, Value: 500.75
	fmt.Println("Value: ", b) // Type: float64, Value: 1.7e308
}

func main() {
	allFloat32()
	allFloat64()

	// this will show error on the x with float32 type
	// constant 3.4e39 overflows float32
	// uncomment below and see the error
	// var x float32 = 3.4e39
	// fmt.Println(x)
}
