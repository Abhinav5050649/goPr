package main

import "fmt"

type myAlias = string
type myDefined string

func main(){
	/*
	var str myAlias= "I am an alias"
	fmt.Println("%T - %s", str, str)

	var strr myDefined = "I am defined"
	fmt.Printf("%T - %s", strr, strr)
	*/

	var alias myAlias
	var def myDefined

	var copy1 string = alias

	var copy2 string = def

	fmt.Println(copy1, copy2)
}

