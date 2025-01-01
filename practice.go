package main
import "fmt"
func formatSpecifier(){
	var a float64 = 45.98
	b:= 45.34
	var name string = "Kamaljeet"
	var exp int = 5
	//printf 
	//`` are used to print using multiple lines
	fmt.Printf(`value of A : %v
	value of B : %v
	`,a,b)
	//sprintf use case: instead of writing the output to the console, it returns the formatted string for further use in your program.
	result := fmt.Sprintf("My name is %v and I have %d years of experience.\n",name,exp)
	fmt.Println(result)
}

func doubleReturn(text1, text2 string)(name string ,age string ){
	name = text1
	age = text2
	return name, age		//else we can simply write "return" to return all values
}