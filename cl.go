package main

import "fmt"

func main(){
	/*
	x := 10

	if x > 5{
		fmt.Println("x is gt 5")
	}else if x > 10{
		fmt.Println("x is gt 10")
	}else{
		fmt.Println("else case")
	}
	*//*
	day := "monday"

	switch day{
	case "monday":
		fmt.Println("time to work!")
	case "friday":
		fmt.Println("let's party!")
	default:
		fmt.Println("Do stuff!")
	}*/

	x :=  10

	switch {
	case x > 5:	
		fmt.Println("x gt 5")
	default:
		fmt.Println("x is !gt 5")
	}

}
