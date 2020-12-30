package main

import "fmt"

type User struct {
	name          string
	monthlySalary int64
	time          int64
}

var userArr [3]User

func main() {
	var name string
	var money, time int64
	// Initializing [3]Users
	for i := range userArr {
		fmt.Scan(&name)
		userArr[i].name = name
		fmt.Scan(&money)
		userArr[i].monthlySalary = money
		fmt.Scan(&time)
		userArr[i].time = time
	}
	//fmt.Println(userArr)
	// Output [3]Users
	fmt.Println("№ | Name | Payment | Works | TotalPayment")
	for i := range userArr {
		//tmp := int64(userArr[i].monthlySalary*userArr[i].time)
		//fmt.Printf("%T %d", tmp , tmp)
		fmt.Printf("%d | %4s | %7d | %5d | %8d \n", i, userArr[i].name, userArr[i].monthlySalary, userArr[i].time,userArr[i].monthlySalary*userArr[i].time )
	}
}