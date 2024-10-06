package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "résumé"
	fmt.Println(myString)
	var indexed = myString[0]
	fmt.Println(indexed)
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}
	fmt.Printf("\nThe length of myString is %v", len(myString)) //Size in bytes

	//Easier way to iterating an indexed string is to cast them to Array or Rune
	var myString2 = []rune("résumé")
	fmt.Printf("\nThe length of myString is %v\n", len(myString2)) //Length of the string
	var indexed2 = myString2[1]
	fmt.Printf("%v, %T\n", indexed2, indexed2)

	//Now when iterating a rune we will get continious index
	for i, v := range myString2 {
		fmt.Println(i, v)
	}

	//Rune type can also be declared by using single quote
	var myRune = 'a'
	fmt.Printf("\nmyRune = %v", myRune)

	var strSlice = []string{"s", "u", "p", "e", "r"}
	var catStr = ""

	for i := range strSlice {
		catStr += strSlice[i]
	}
	fmt.Printf("\n%v", catStr)

	//Strings in go are immutable
	//Below approach is much more efficient because it only creates the string at the end
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr2 = strBuilder.String()
	fmt.Printf("\n%v", catStr2)
}
