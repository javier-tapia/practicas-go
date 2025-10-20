package main

import (
	"fmt"
	"math"
	"strconv"

	"rsc.io/quote"
)

// Declaración de constantes
const Pi float32 = 3.14 // También se puede omitir el tipo: 'const Pi = 3.14'

const (
	x = 100
	y = 0b1010 // Binario
	z = 0o12   // Octal
	w = 0xFF   // Hexadecimal
)

const (
	Domingo   = iota + 1 // 1
	Lunes                // 2
	Martes               // 3
	Miercoles            // 4
	Jueves               // 5
	Viernes              // 6
	Sabado               // 7
)

func main() {
	fmt.Println("Hello world from Javi")
	fmt.Println(quote.Hello())

	// Declaración de variables.
	// Con var, se pueden declarar dentro o fuera de una función
	// var firstName, lastName, age = "Javi", "Tapia", 40

	// Esta forma con ':=' (y omitiendo el 'var'), sólo se puede usar dentro de una función
	firstName, lastName, age := "Javi", "Tapia", 40

	fmt.Println(firstName, lastName, age)
	fmt.Println(Pi)
	fmt.Println(x, y, z, w)
	fmt.Println(Viernes)
	fmt.Println(math.MinInt8, math.MaxInt8)
	fmt.Println(math.MaxUint8)

	fullName := "Javi Tapia \t(alias \"javito\")\n"
	fmt.Println(fullName)

	var a byte = 'a'
	fmt.Println(a) // 97 -> Valor decimal para la letra 'a' en código ASCII

	s := "Hola"
	fmt.Println(s[0]) // 72 -> Valor decimal para la letra 'H' en código ASCII

	var r rune = '👍'
	fmt.Println(r) // 128077 -> Valor decimal para el caracter Unicode

	// Conversión de tipos
	var integer16 int16 = 50
	var integer32 int32 = 100
	fmt.Println(int32(integer16) + integer32) // 150

	numberString := "100"
	i, _ := strconv.Atoi(numberString)
	fmt.Println(i + i) // 200

	n := 42
	s = strconv.Itoa(n)
	fmt.Println(s + s) // 4242

	// Formatear un string
	fmt.Printf("Hola, soy %s y tengo %d años\n", firstName, age)

	greeting := fmt.Sprintf("Hola, soy %s y tengo %d años", firstName, age)
	fmt.Println(greeting)

	// Solicitar datos de entrada al usuario
	// fmt.Print("Ingrese su nombre: ")
	// fmt.Scanln(&firstName)
	// fmt.Print("Ingrese su edad: ")
	// fmt.Scanln(&age)
	// fmt.Printf("Hola, soy %s y tengo %d años\n", firstName, age)

	// Arrays
	var arrayOne [5]int
	arrayOne[0] = 10
	arrayOne[1] = 20
	var arrayTwo = [5]int{10, 20, 30, 40, 50}
	arrayTwo[4] = 60
	fmt.Println(arrayOne, arrayTwo)
	var arrayThree = [...]int{11, 22, 33, 44, 55}
	for i, v := range arrayThree {
		fmt.Println(i, v)
	}
	var arrayFour = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(arrayFour)

	// Slice
	weekDays := []string{"Domingo", "Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado"}
	fmt.Println(weekDays)
	sliceWeekDays := weekDays[0:5]
	fmt.Println(sliceWeekDays)
	fmt.Println(cap(sliceWeekDays)) // 7 -> Capacidad del slice original, no de lo que se está tomando (0:5)
	sliceWeekDays = append(sliceWeekDays, "Osvaldo")
	fmt.Println(sliceWeekDays) // [Domingo Lunes Martes Miercoles Jueves Osvaldo]
	sliceTwo := []string{"Urano", "Plutón"}
	sliceWeekDays = append(sliceWeekDays, sliceTwo...)
	fmt.Println(sliceWeekDays) // [Domingo Lunes Martes Miercoles Jueves Osvaldo Urano Plutón]
	sliceWeekDaysDeleteJueves := append(sliceWeekDays[:4], sliceWeekDays[5:]...)
	fmt.Println(sliceWeekDaysDeleteJueves) // [Domingo Lunes Martes Miercoles Osvaldo Urano Plutón]

	// Punteros
	var x int = 10
	var p *int = &x
	fmt.Println(x) // 10
	fmt.Println(p) // 0x1400010e4f0
	editar(p)
	fmt.Println(x) // 20
	person := Person{"Javi", 40, "javi@gmail.com"}
	person.sayHello()
}

func editar(x *int) {
	*x = 20
}

func (p *Person) sayHello() {
	fmt.Println("Hi, my name is", p.name) // Hi, my name is Javi
}

type Person struct {
	name  string
	age   int
	email string
}
