package main

import (
	"fmt"
	"math"
)

func sumOfSinXY(x, y float64) float64 {
	return (math.Sin(float64(x)) + math.Sin(float64(y)))
}

func main() {
	var firstNum, secondNum float64
	fmt.Println("Input first & second numbers")
	fmt.Scan(&firstNum, &secondNum)
	result := sumOfSinXY(firstNum, secondNum)
	fmt.Printf("Sin(%.2f) + Sin(%.2f) = %f", firstNum, secondNum, result)
}
