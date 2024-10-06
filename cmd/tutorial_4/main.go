package main

import "fmt"

func main() {
	// Arrays - is a fixed length collection of data all of the same type, indexable
	// and stored in a conigious memory locations
	fmt.Println("Array")

	var intArr [3]int32
	fmt.Println(intArr[0])
	intArr[2] = 2
	fmt.Println(intArr[1:3])

	// int32 is 4 bytes of memory and array is of 3 size, 12 bytes of contigious memory are allocated to this array
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])
	// The & before the variable name prints out memory location of the variable

	var intArr2 [3]int32 = [3]int32{1, 2, 3}
	intArr3 := [3]int32{1, 2, 3}
	intArr4 := [...]int32{1, 2, 3}
	fmt.Println(intArr2)
	fmt.Println(intArr3)
	fmt.Println(intArr4)

	// Slices - Slices are just wrappers around arrays
	fmt.Println("Slice")
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7) //Unlike arrays, value can be added to slice
	fmt.Println(intSlice)
	fmt.Printf("\nThe length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...) // Spread operator ...
	fmt.Println(intSlice)

	// Another way of creating a slice is using make function
	var intSlice3 []int32 = make([]int32, 3)
	fmt.Println(intSlice3)

	// Map - is a set of key value pairs where you can look up a value by its key
	var myMap map[string]uint8 = make(map[string]uint8)
	// this means keys are of type string and values are unsigned int8
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 21, "Sarah": 18}
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Jason"]) // Returns default value
	var age, ok = myMap2["Jason"]
	delete(myMap2, "Adam")
	if ok {
		fmt.Printf("\nThe age is %v\n", age)
	} else {
		fmt.Println("Invalid name")
	}

	//Loops
	fmt.Println("Loopig Map")
	for name := range myMap2 {
		fmt.Printf("\nName %v ", name)
	}
	// when iterating over a map, order is not preserved
	myMap2["Carlos"] = 45

	fmt.Printf("\nLoopig Map with value")
	for name, age := range myMap2 {
		fmt.Printf("\nName %v , Age:%v", name, age)
	}

	fmt.Printf("\nLoopig Array with index and value")
	for i, v := range intArr {
		fmt.Printf("\nIndex %v , Value:%v", i, v)
	}

	fmt.Printf("\nLoopig Slice with index and value")
	for i, v := range intSlice {
		fmt.Printf("\nIndex %v , Value:%v", i, v)
	}

	// Go does not have while loops, but it can be achieved with for
	fmt.Println("\nWhile loop")
	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("While loop with break")
	i = 0
	// It can also be achieved with break in an if statement inside the loop
	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("Usual loop")
	//Usual for loop
	for i = 0; i < 10; i++ {
		fmt.Println(i)
	}

}
