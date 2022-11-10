package main

import "fmt"

/*
func myNextFunction(p1 string) string{
	msg := fmt.Sprintf("%s function", p1)
	return msg
}
*/
/*
func myFunc(p1 string) (string, int){
	msg := fmt.Sprintf("%s function", p1)
	return msg, 10
}
*/
/*
func myFunc(p1 string) (s string, i int){
/*	s = fmt.Sprintf("%s function", p1)
	i = 10

	return
}
*/
/*
func myFunc(){
	fn := func(){
		fmt.Println("inside fn")
	}

	fn()
}
fu
*/

func add(values ...int) int {
	sum := 0

	for _, v := range values{
		sum += v
	}

	return sum
}
func main(){
	s:= add(1, 2, 3)
	fmt.Println(s)
}
