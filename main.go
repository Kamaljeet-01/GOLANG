package main

import (
	"fmt"
	emp "golang/subFolder1"
	//importing 3rd party package
	"github.com/grandper/go-randomdata"
)

// it will give you error if "we declare variable but don't use"
func main() {
	r := randomdata.FromSeed(1234)
	fmt.Println("Hello Kamaljeet !! ")
	// var num1 int
	// fmt.Scan(&num1)
	// var num2 int
	// fmt.Scan(&num2)
	// fmt.Println(num1+num2)
	//fmt.Println("Enter details :")
	// fmt.Print(intrestCalculator())
	//emp.DisplayDetails("Kamaljeet", 75000)

	//Basic if-else switch cases :
	emp.Basics()
	formatSpecifier()
	name, age := doubleReturn("Kamaljeet","21")
	fmt.Printf("Name : %v and age : %v\n",name,age)
	emp.WriteBalanceToFile("My name is kamaljeet.\nThis is demo file.")
	fmt.Println(emp.ReadFromFile())
	//getting random number from 0 to 400 using third party package
	fmt.Println(r.Number(400))
}

