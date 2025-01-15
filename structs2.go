package main

import "fmt"

//defining struct
type Circle struct {
	radius float64
}

//creating circle's perameter method
func (r Circle) areaOfCircle() float64 {
	return r.radius * 2
}

//We wrote this function to do the same task as we do in main.go file while calling that area function in main function.
//Its basically an indirect approach
func structs() {
	
	var rad float64
	fmt.Println("Enter radius : ")
	fmt.Scan(&rad)

	//populating struct and making instance from this circle struct
	var gola Circle
	gola = Circle{
		radius: rad,
	}
	fmt.Println("Area of circle is  ", gola.areaOfCircle())
}


