package main

import "fmt"

func main(){

	/*
	a:=10;
	var p *int = &a;
	fmt.Println("Address: ", p);
	fmt.Println("Value: ", *p);
	*/
	/*
	a := 10;
	var p *int = &a;
	fmt.Println("before: ", a);
	fmt.Println("Address: ", p);

	*p = 20;

	fmt.Println("After: ", a);
	*/

	p:= new(int);
	*p = 10;

	p1 := &p;

	fmt.Println("P Value: ", *p, " Address: ", p);
	fmt.Println("P1 Value: ", *p1, " Address: ", p);

	fmt.Println("Dereferenced Value: ", **p1);
}	
