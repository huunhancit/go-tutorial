package main

import (
	"fmt"
	"math"
	"math/rand"
)

func sqrtMath() {
	fmt.Println("Sqrt 4 is", math.Sqrt(4))
}

func randMath() {
	fmt.Println("A number 1 - 100: ", rand.Intn(100))
}

func main() {
	sqrtMath()
	randMath()
}
