package main
import "fmt"

func main(){
	/*
	for i := 1; i <= 10; i++	{
		fmt.Print(i, " ")
	}
	fmt.Println()*/

	for i:= 0; i <= 10; i++ {
		if i == 0 {
			continue
		}

		fmt.Print(i, " ")

		if i > 9 {
			break
		}
	}

	fmt.Println()

	/*for {
		fmt.Println("hello");
	}*/
}
