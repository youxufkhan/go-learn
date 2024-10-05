package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	// var result float32 = floatNum32 + intNum32 // Not possible because adding numbers arent the same type
	var result float32 = floatNum32 + float32(intNum32) //  this will work
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2) // Integer division results in integer and the result is rounded down
	fmt.Println(intNum1 % intNum2) // can use % to print remainder

	var myString string = "Hello \nWorld"
	var myString2 string = `Hello
	World` //backqoutes support string formatting
	fmt.Println(myString)
	fmt.Println(myString2)

	var concatedString string = "Hello" + " " + "World"
	fmt.Println(concatedString)

	fmt.Println(len("test")) // Will print the bytes of the string not the length
	fmt.Println(len("γ"))    // Will print the bytes of the string not the length
	fmt.Println(utf8.RuneCountInString("γ"))

	var myBoolean bool = false
	fmt.Println(myBoolean)

	var intNum3 int      // go will set default value
	fmt.Println(intNum3) //default value is 0

	var myVar = "text" // We can omit the type when creating variable, in this case the type is inferred
	myVar2 := ":text"  // We can go further and drop the var keyword and use the shorthand :=
	fmt.Println(myVar)
	fmt.Println(myVar2)

	var var1, var2 int = 1, 2
	var var3, var4 int = 3, 4
	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(var3)
	fmt.Println(var4)

	const myConst string = "const value"
	// myConst = "another const value" // this cant be changed
	// const myConst string // also constants cant just be declared and assigned later
	fmt.Println(myConst)

}
