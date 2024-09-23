package main

import (
	"bufio"
	"fmt"
	"os"
)
//it will give you error if "we declare variable but don't use"
func main() {
	fmt.Println("Hello Kamaljeet !! ")
	reader := bufio.NewReader(os.Stdin)
	var num, err = reader.ReadString('\n')	//this reader.ReadString gives two values: one input and second is if input is not valid or something than error msg.
	fmt.Println("Number u entered is : ", num)
	fmt.Println("Error occured while running program :", err)//Printing Error msg

}
