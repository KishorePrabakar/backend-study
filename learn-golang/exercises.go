package main

import ("fmt")

func main(){
	/*
// 1. Variable declaration

// Type-1: Explicit type casting
var x int = 5
var name string = "Kishore"

fmt.Println("Hello, " + name)
fmt.Println(x)
fmt.Println()

// Type-2 : Type Inference
y := 25
place := "Chennai"
sum := x + y

fmt.Println(name,", you are from", place)
fmt.Println(x,y)
fmt.Println(sum)

// Type - 3: Constants

const val int = 10
var obj int = 5
// val = val + 1 (produces error -  cannot assign to val (neither addressable nor a map index expression))
obj = obj + 5
fmt.Println(val)
fmt.Println(obj)
*/

// 2. If / else
/* var age int = 18
if age >= 18{
	fmt.Println("Eligible to vote")
}else if age < 18{
	fmt.Println("Not Eligible to vote")
}
*/

// 3. For loop

/*for i := 1; i < 6; i++{
	fmt.Println(i)
}

//while loop equivalent
*/
var n int
for n <= 100{
	fmt.Println(n)
	n+=10
}


}


