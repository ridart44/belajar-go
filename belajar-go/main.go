package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Hello", "World", "Test") // printing in Go

	/* var text string = "Hello" // store var in declared data type
	secondVar := 3 */ // decalre variable without define it's data type
	/* var firstVar, secondVar = "Hello", 4 // multiple var declared datatypes

	thirdVar, forthVar := "Hello", "World" // multiple var undeclared datatypes
	_ = 5 // uderscore as null is used to store unused variable
	fmt.Println(firstVar, secondVar) // combine printing var
	fmt.Println(thirdVar, forthVar) */

	/* dummy := new(string) // new is used to declared pointer variable with a defined data types
	fmt.Println(dummy) //
	fmt.Println(*dummy) */
	// oh i need to relearn C

	//-----------------------------Data Types-----------------------------
	//Numeric non-decimal
	// int and uint
	// int8, int32, uint16,uint8, etc.
	// int is signed int  -- formatting %d

	// Numeric decimal
	// float32, float64 -- formatting %f

	// Boolean
	// bool true, bool false -- formatting %t

	// string
	// use "" or `` -- `` will define all char inside as a string

	// nil & Zero value
	// all data types has zero value as default value.
	// zero value of string is "" (empty string)
	// 				 bool   is false
	//numeric non decimal   is 0
	//    numeric decimal   is 0.0
	// nil and null are same essentially
	// data types that can be setted as nil:
	// pointer, function data types, slice, map, channel

	//-----------------------------Constant-----------------------------
	/* const firstName string = "Hary"
	fmt.Println(firstName)
	var lastName = "Ridart"
	fmt.Println("Hello", firstName, lastName) */

	//-----------------------------Operator-----------------------------
	// arithmetic, comparator, operand are just like C
	// + - * / % == != < <= > >= && || !
	/* left := true
	right := false
	fmt.Println(left || right)
	var1 := 5
	var2 := "s"
	fmt.Println(reflect.TypeOf(var1) == reflect.TypeOf(var2)) */

	//-----------------------------Selection-----------------------------
	// if, else, else if
	a := 5
	b := 7
	c := 4
	if a > c {
		fmt.Printf("%d adalah a, %d adalah b, %d adalah c", a, b, c)
	} else if a < c {
		fmt.Println("Condition else if")
	} else {
		fmt.Println("Else condition")
	}

}
