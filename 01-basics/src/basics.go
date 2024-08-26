package main

import (
	"fmt"
	"reflect"
)

// if class name starts letter is minus the class is private, if mayus is public
type car struct {
	brand string
	year  int
}

func yes() (text string) {
	// The defer statement makes sure that whatever function is assigned to it gets executed for sure even if the function panics or the code returns well before the end of the function.
	defer func() {
		text = "no"
	}()
	return "yes"
}

func main() {
	fmt.Println("hello world!")
	//const
	const pi float64 = 3.14
	const pi2 = 3.141592653589793238462643383279502884197

	fmt.Println("pi", reflect.TypeOf(pi), "is", pi)
	fmt.Println("pi2", reflect.TypeOf(pi2), "is", pi2)
	//vars
	b := 12
	var h int = 14
	var a int

	fmt.Println("base:", b, "height:", h, "area:", a)

	//zero variables
	var aa int     //0
	var bb float64 //0
	var cc string  //''
	var dd bool    //false
	fmt.Println(aa, bb, cc, dd)

	// square area
	const squareB = 10
	areaSquare := squareB * squareB
	fmt.Println("Area:", areaSquare)

	// imaginary
	c := 10 + 8i
	fmt.Printf("Imaginary is %v\n", c)
	fmt.Printf("Imaginary is %T\n", c)

	//for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//for while
	counter := 10
	for counter < 20 {
		fmt.Println(counter)
		counter++
		continue
		fmt.Println("ayayay")
	}

	//Defer
	fmt.Println(yes())

	//Arrays and slices
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	slice := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(slice)

	//Map
	m := make(map[string]int)
	m["jose"] = 14
	m["pepe"] = 13
	m["pepe"] = 23

	fmt.Println(m)
	for i, v := range m {
		fmt.Println(i, v)
	}

	//pointers
	var creature string = "shark"
	var pointer *string = &creature

	fmt.Println("creature =", creature)
	fmt.Println("pointer =", pointer)

	//Numeros enteros
	//int = Depende del OS (32 o 64 bits)
	//int8 = 8bits = -128 a 127
	//int16 = 16bits = -2^15 a 2^15-1 = 32768
	//int32 = 32bits = -2^31 a 2^31-1 = 2147483648
	//int64 = 64bits = -2^63 a 2^63-1 = 9.223372e+18

	//Optimizar memoria cuando sabemos que el dato siempre va ser positivo
	//uint = Depende del OS (32 o 64 bits)
	//uint8 = 8bits = 0 a 127
	//uint16 = 16bits = 0 a 2^15-1
	//uint32 = 32bits = 0 a 2^31-1
	//uint64 = 64bits = 0 a 2^63-1

	//numeros decimales
	// float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
	// float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308

	//textos y booleanos
	//string = ""
	//bool = true or false

	//numeros complejos
	//Complex64 = Real e Imaginario float32
	//Complex128 = Real e Imaginario float64
	//Ejemplo : c:=10 + 8i

}
