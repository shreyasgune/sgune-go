package main

import (
	"fmt"
	// "strings"
	"sync"
	"time"
	"standardref/helper"
)


func basicDataTypeRef() {

	// When to use int types

	/* For Big Numbers that are whole integers
	-2147483648 to 2147483647 in 32 bit systems and -9223372036854775808 to 9223372036854775807 in 64 bit systems
	*/
	var intType int = 10;

	/*for regular numbers that don't go beyond -128 to 127 | takes 1 byte*/
	var int8Type int8 = 126;

	/* For numbers between -32768 to 32767 | takes 2 bytes */
	var int16Type int16 = 32766;

	/* For -2147483648 to 2147483647 | takes 4 bytes */
	var int32Type int32 = 2147483646;

	/* non integer BIG numbers */
	var float32Type float32 = 1.0; // 4 bytes | -3.4e+38 to 3.4e+38.
	var float64Type float64 = 1.1; // 8 bytes | -1.7e+308 to +1.7e+308.

	// is it TRUE or is it FALSE?
	var trueOrFalse bool = false;

	// Strings
	var name string = "Obi Wan Kenobi"

	fmt.Println(intType, int8Type, int16Type, int32Type, float32Type, float64Type, trueOrFalse, name)

	// Complex Numbers
	var complexNum64 complex64 = 1 + 2i
	var complexNum128 complex128 = 1 + 3i
	fmt.Println("Complex Numbers:", complexNum64, complexNum128)

	// RAW BYTES - to handle binary files or I/O Operations
	var b byte = 'A'
	fmt.Println("Raw bytes for A char", b)

	//Special character, used for unicode
	var r rune = '界'
	fmt.Println("Print rune: ", r)

}

func compositeDataTypeRef() {
	// when size is fixed in advance, use: array
	var arr [5]int = [5]int{1,2,3,4,5}
	fmt.Println("Array of length", len(arr) , "is:", arr)

	// dynamic array is basically use: slice
	var s []int = []int{1,2,3,4,5,6}
	s = append(s,7)
	fmt.Println("Slice of length", len(s) , "is:", s)

	// key value pairs, like dicts : use maps
	m := map[string]int {
		"Alice" : 22,
		"Bob"	: 23,
	}
	fmt.Println("Map:",m)

	// when you have a custom data type: use struct
	type Person struct {
		Name string
		Age  uint8
	}
	person := Person {
		Name: "Alice",
		Age: 25,
	}
	fmt.Println("Struct:",person) 
}


//Pointers
//Pointer : holds memory address of a var, use for efficient memory access or modifying variables through references.

func increment(x *int) { 
	//since we're getting the address of x, we gotta use pointers to extract the value at that address.
	*x =*x + 1;
}
func pointerRef() {
	a := 10;
	fmt.Println("Before increment", a)
	increment(&a) //pass the address of 'a'
	fmt.Println("After increment", a)
}


// Interface
// When you don't know what the return type of your function is going to be, basically when you want to pass any datatype.
//They enable polymorphism and decoupling in your code.

func describe(i interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", i, i)
}
type Shape interface {
	Area() float32
}
type Circle struct {
	Radius float32
}
func (c Circle) Area() float32 {
	return 3.14 * c.Radius * c.Radius
}
func printArea(s Shape) {
	fmt.Println("Area:", s.Area())
}
func AreaFinder() {
	/*
		The Shape interface requires an Area() method.
		Circle implements Shape by defining the Area() method.
		printArea works with any type that satisfies Shape interface.
	*/

	c := Circle{Radius: 5}
	printArea(c)
}


// Special Type
// function, you know, to define a function which shares your attributes (self explanatory)

// chan : safe communication between go routines
func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Println("Sending:", i)
		ch <- i // Send data to the channel
	}
	close(ch) // Close the channel when done
}
func chanComms() {
	ch := make(chan int) //create a channel

	//Goroutine sends data
	go producer(ch)
	for value := range ch {
		fmt.Println("Received:", value)
	}
	fmt.Println("Channel communication complete.")	
}



// Custom Type
func customType() {
	type Age int

	var myAge Age = 30
	fmt.Println("My Age:", myAge)
}


// Go routines - run on separate threads, but share the mem space
// lets simulate some work and also, for loops
func printNumbers(n int) {
	for i := 1; i<=n; i++ {
		fmt.Println("Number:",i)
		time.Sleep(100* time.Millisecond) //Timepass
	}
	wg.Done()
}

// Taking user input
func takeUserInput() {
	var theName string	
	fmt.Println("Enter name:")
	fmt.Scan(&theName)
	fmt.Println("You entered",theName)
}

// Loops and Conditionals
func loopAndCondish(x int) {
	for {
		if x >10 {
			continue
		} else {
			break
		}
	}

	switch x {
		case 1:
			fmt.Println("You entered 1")
		default:
			fmt.Println("You did not enter 1")
	}
}

var wg = sync.WaitGroup{}

func main() {
	fmt.Println("This is the main function")

	fmt.Println("\n\nBasic Data Type Ref")
	basicDataTypeRef()

	fmt.Println("\n\nComposite Data Type Ref")
	compositeDataTypeRef()

	fmt.Println("\n\nInterface")
	describe("hello")
	describe(42)
	describe(3.14)
	AreaFinder()

	fmt.Println("\n\nChannel Comms")
	chanComms()

	fmt.Println("\n\nCustom Types: type Age int ==")
	customType()

	fmt.Println("\n\nUser Input:")
	takeUserInput()

	fmt.Println("\n\nLoops and Condish")
	loopAndCondish(1)

	fmt.Println("\n\nHelper stuff")
	helper.JustPrint()


	fmt.Println("\n\nGoroutine stuff")
	wg.Add(1)
	go printNumbers(5)
	fmt.Println("GoRoutine started")
	time.Sleep(1 * time.Second)
	wg.Wait()
	fmt.Println("MAIN IS OVER")

}