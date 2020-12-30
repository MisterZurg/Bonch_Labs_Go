package main

import "fmt"

func main() {
	var x, y int64
	var op string
	fmt.Scan(&x, &op, &y)
	switch op {
	case "+":
		fmt.Println(add(x, y))
	case "-":
		fmt.Println(sub(x, y))
	case "*":
		fmt.Println(mul(x, y))
	case "/":
		/* In case we need to continue using our calc
		   defer func() {
		      if r:= recover(); r != nil {
		         fmt.Println("recovered from function with divide by zero")
		      }
		   }()
		*/
		fmt.Println(div(x, y))
	case "%":
		/* In case we need to continue using our calc
		   defer func() {
		      if r:= recover(); r != nil {
		         fmt.Println("recovered from function with divide by zero")
		      }
		   }()
		*/
		fmt.Println(rem(x,y))
	}
}
func add(first, second int64) int64 {
	return first + second
}
func sub(first, second int64) int64 {
	return first - second
}
func mul(first, second int64) int64 {
	return first * second
}
func div(first, second int64) int64 {
	if second == 0{
		panic("division by zero!")
	}
	return first / second
}
func rem(first, second int64) int64 {
	// Throw panic if there is a division by zero
	if second == 0{
		panic("division by zero!")
	}
	return first % second
}
