package main

import "fmt"

// Struct is the way of defining your own type

type gasEngine struct {
	mpg       uint8
	gallons   uint8
	ownerInfo owner
}

type owner struct {
	name string
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

type engine interface {
	milesLeft() uint8
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it")
	} else {
		fmt.Println("You cant make it")
	}
}

func main() {
	// Struct and Types
	var myEngine gasEngine
	fmt.Println(myEngine.mpg, myEngine.gallons)

	var myEngine2 gasEngine = gasEngine{mpg: 25, gallons: 15}
	fmt.Println(myEngine2.mpg, myEngine2.gallons)

	var myEngine3 gasEngine = gasEngine{28, 17, owner{"Alex"}}
	fmt.Println(myEngine3.mpg, myEngine3.gallons)

	myEngine3.mpg = 20
	fmt.Println(myEngine3.mpg, myEngine3.gallons, myEngine3.ownerInfo.name)

	//Anonymous struct

	var myEngine4 = struct {
		mpg     uint8
		gallons uint8
	}{33, 100}
	fmt.Println(myEngine4.mpg, myEngine4.gallons)

	fmt.Printf("Total miles left in tank: %v", myEngine2.milesLeft())

	var myEngine5 electricEngine = electricEngine{25, 15}
	canMakeIt(myEngine5, 50)

}
