package main

import (
	"fmt"
	"log"
	"time"
)

func main()  {
	// Do smth
}

func mutaluzm(myPassword string )  {
	type brokenAuthenticationForm struct {
		login string
		password string
	}
	mybrokenAuthenticationForm := brokenAuthenticationForm{
		login:"admin",password:"su",
	}
	// Тут даже не совсем инъекция, а поиск "бага"					   \/
	if mybrokenAuthenticationForm.password == myPassword || myPassword == "omega"{
		fmt.Println("Access approved")
		fmt.Println("alert(1)") // При возможности XXE || RCE можно закрепиться в системе
	}else{
		fmt.Println("Access denied")
	}
}

func commensaluzm(){
	type brokenAuthenticationForm struct {
		login string
		password string
	}
	mybrokenAuthenticationForm := brokenAuthenticationForm{
		login:"admin",password:"su",
	}
	// Zero-day допустим тестировщики не проверили,
	// что будет если отправить пустой пароль
	if mybrokenAuthenticationForm.password == ""{
		fmt.Println("Access approved")
		fmt.Println("Execute order 66")	// Неожиданное поведение программы,
											// поведение из-за spoiled code
		executeOrder66();  // Spiled DI
	}
}
func executeOrder66(){
	for i:=0;i<66;i++{
		fmt.Println("Execute order 66")
	}
}
func allelopatia(){
	// Logic Bomb
	// Закрепившись в системе можно, с помощью логической бомбы
	// совершить вредоносную активность при этом положить алгоритм
	// или всю систему
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	fmt.Println(then)

	if then.Equal(time.Now()){
		// delete(program_name.go)
		log.Fatal()
	}
}