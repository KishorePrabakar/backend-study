package main

import (
	"fmt"
	// "math"
)

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

var n int
for n <= 100{
	fmt.Println(n)
	n+=10
}

// 4. Functions
a := 5.5
var b float64 = 4
c := add(a,b)
fmt.Println(c)

var a int = 3
var a_sqr float64 = square(a)
fmt.Println(a_sqr)

//call by reference
var b float64 = 3
fmt.Println(b)
var b_cube float64 = cube(&b)	
fmt.Println(b_cube)
fmt.Println(b)



var a int = 17
var b int = 3
fmt.Println(a/b)
fmt.Println(a%b)

var x string = "Kishore "
var y string = "Prabakar"
var z string = x + y
fmt.Println(z)


var x int = 10
var y int = 12

if x > 5 && y == 10{
	fmt.Println(true)
}else{
	fmt.Println(false)
}

var a1 string = "apple"
var b1 string = "banana"

if(a1 < b1){
	fmt.Println("large")
}else{
	fmt.Println("small")
}

fmt.Println(4 & 5)


var x int = 5
var y int64 = 10
z := x + int(y)
fmt.Println(z)

var a int = 100
var b *int = &a
fmt.Println(&a)
fmt.Println(b)
*/

fmt.Println(reverse("Kishore 😎😎"))

}

func reverse(a string) string{
	runes := []rune(a)
	for i, j := 0, len(runes) - 1; i < j; i, j = i+1, j-1{
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)

}

/*func cube(b *float64) float64{
	*b = math.Pow(*b, 3)
	return *b
}

func square(a int) float64{
	return math.Pow(float64(a),2)
}

func add(a float64, b float64) float64{
	return a + b
}
*/
