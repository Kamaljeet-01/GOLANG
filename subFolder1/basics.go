package emp

import (
	//"errors"
	"fmt"
	"os"
)

//if else syntax
func Basics(){
	var a, b int = 3, 5
	if a > b {
		fmt.Println("a is greater than b")
	} else {
		fmt.Println("b is greater than a")
	}

	//For loop syntax
	for c:=5;c>0;c-- {
		fmt.Println(c)
	}

	//Switch statement syntax
	fmt.Println("Switch Statement :")
	switch (a){
	case 1:
		fmt.Println("a is 1")
	case 2:
		fmt.Println("a is 2")
	default:
		fmt.Println("a is not 1 or 2")
	}
}

func WriteBalanceToFile(balance string){
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText),0644) 
}

func ReadFromFile() (string, error){
	data, err:=os.ReadFile("balance.txt")
	content := string(data)
	if(err != nil){
		//fmt.Println("Got an error !!! :",err)
		//return content, errors.New("error reading file")
		//or we can use panic function
		panic("can't continue this program.Sorry !!")
	}
	
	return content,err
}