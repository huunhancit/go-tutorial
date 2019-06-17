package main

import "fmt"

func add(x float64, y float64) float64 {
	return x + y
}

func multiple(a, b string) (string, string, string) {
	return a, b, a + b
}

func main() {
	var num1 float64 = 1.0
	fmt.Println("Result : ", add(1.2, 2.2))
	fmt.Println("Numberone: ", num1)
}
