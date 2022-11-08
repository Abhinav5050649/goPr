package main

import "fmt"

func main(){
	var foo string = "Hello World"
	fmt.Println(foo)

	var s1, s2 = "hello" , "world"

	fmt.Println(s1 + s2)
	fmt.Println(s1, s2)

	const constant = "This is a constant"

	fmt.Println(constant)

	var bio = `This is a statically typed.
			I was designed at Google`

	fmt.Println(bio)

	var choice, choi1 = true, false

	fmt.Println(choice, choi1)

	
	var i int = 404;
	var i8 int8 = 127;
	var i16 int16 = 32767;
	var i32 int32 = -217483647;
	var i64 int64 = 9027917295909258;

	fmt.Println(i, i8, i16, i32, i64);
		
	var ui64 uint64 = 17329890545;

	fmt.Println(ui64);
	
	//type byte = uint8;
	//type rune = int32;

	var b byte = 'a';

	fmt.Println(b);
	
	var f32 float32 = 1.7812;
	var f64 float64 = 3.1415;

	fmt.Println(f32, f64);

	var c1, c2 = complex(10, 1), 12 + 4i;
	fmt.Println(c1, c2);
	
	/*
	var ih int;
	var f float64;
	var bh bool;
	var s string;

	fmt.Println(ih, f, bh, s);
	*/

	ih := 42;
	f := float64(ih);
	u := uint(f);

	fmt.Println(ih, f, u);

	fmt.Printf("%T %T", f, u);
	fmt.Println();
}
