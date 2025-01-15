package main

import "fmt"

//Rectangle struct
type Rect struct {
	len int
	bre int
}

//circle struct
type Circle2 struct {
	rad float64
}

//function for calculating area of rectangle
func (r Rect) AreaOfRect() int {
	return (r.len * r.bre)
}

//function to calculate area of circle
func (c Circle2) AreaOfCircle() float64 {
	return (3.14 * c.rad * c.rad)
}

//Made this function so that i can create an instance for both struct and populate them in this function and call this function in main file
func AreafromStruct() {
	var length, breadth int
	var radius float64
	fmt.Println("Enter Length : ")
	fmt.Scan(&length)
	fmt.Println("Enter Breadth : ")
	fmt.Scan(&breadth)
	fmt.Println("Enter Radius : ")
	fmt.Scan(&radius)
	//populating struct
	r := Rect{
		len: length,
		bre: breadth,
	}
	c := Circle2{
		rad: radius,
	}

	//calling area method embedded with their respective structs
	fmt.Println("Area of Reactangle : ", r.AreaOfRect())
	fmt.Println("Area of Circle : ", c.AreaOfCircle())

}
