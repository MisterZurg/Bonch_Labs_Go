package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func sumOfSinXYObf(x, y float64) float64 {
	return math.Sin(x) + math.Sin(y)
}

func sumOfSinhXY(x, y float64) float64 {
	return math.Sinh(x) + math.Sinh(y)
}

func sumOfCosXY(x, y float64) float64 {
	return math.Cos(x) + math.Cos(y)
}

func sumOfCoshXY(x, y float64) float64 {
	return math.Cosh(x) + math.Cosh(y)
}

func sumOfTanXY(x, y float64) float64 {
	return math.Tan(x) + math.Tan(y)
}

func sumOfСtgXY(x, y float64) float64 {
	return (math.Cos(x) / math.Sin(x)) + (math.Cos(y) / math.Sin(x))
}

func isDebuggerPresent(processName string) bool {
	loadedStuff := [4]string{
		"IDAPro.exe",
		"WinDbg.exe",
		"Ida Pro",
		"WinDbg",
	}
	for _, a := range loadedStuff {
		if a == processName {
			return true
		}
	}
	return false
}
func isLoadedProcess(processName string) bool {
	processSlice := [2]string{
		"IDAPro.exe",
		"WinDbg.exe",
	}
	for _, a := range processSlice {
		if a == processName {
			return true
		}
	}
	return false
}

func destructiveActions() {
	fmt.Printf("Выполняю деструктивные действия\n" +
		"Performing destructive actions\n" +
		"파괴적인 행동 수행")
}

func dewIt(num int) bool {
	return num%2 == 0
}

func main() {
	start := time.Now()					// Step 5 возвращает текущую дату и время
	fmt.Println("You've started", start.Format("02-01-2006 15:04:05"))
	rand.Seed(time.Now().Unix())   		// Для генерации псевдо-сдучайного числа
	var rndNum = rand.Intn(10 + 1) 	// В пределах от 0 до 10
	fmt.Printf("Your token is %d\n", rndNum)
	if dewIt(rndNum) { 					// Step 3
		fmt.Printf("Acess approved\n")
		var firstNum, secondNum float64
		fmt.Println("Input first & second numbers")
		fmt.Scan(&firstNum, &secondNum)
		result := sumOfSinXYObf(firstNum, secondNum)
		result = sumOfCoshXY(firstNum, secondNum)
		result = sumOfCosXY(firstNum, secondNum)
		result = sumOfTanXY(firstNum, secondNum)
		result = sumOfSinXYObf(firstNum, secondNum)
		fmt.Printf("Sin(%.2f) + Sin(%.2f) = %f\n", firstNum, secondNum, result)
		notResult := sumOfСtgXY(firstNum, secondNum)
		sumOfSinhXY(notResult, notResult)
	} else {
		fmt.Printf("Acess denied")
	}
	finish := time.Now()				// возвращает текущую дату и время

	var name string //= ""				// С целью избежать isDebuggerPresent panic()
	if isDebuggerPresent(name) {		// Step 4
		destructiveActions()
	}


	// вычисляет период между текущим моментом finish и заданным временем в прошлом start
	elapsed := finish.Sub(start)
	if elapsed.Seconds() > 1{			// Step 5
		//destructiveActions()
	}
	if isLoadedProcess(name) {			// Step 6
		destructiveActions()
	}

}
