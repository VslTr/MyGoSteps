СТРУКТУРЫ

// Структура представляет собой значение,
// которое строится из других значений разных типов.
/*
Тип структуры объявляется ключевым словом struct, за ко-
торым следуют фигурные скобки. В фигурных скобках опре-
деляются одно или несколько полей: значений, группируе-
мых в структуре. Определение каждого поля размещается в
отдельной строке и состоит из имени поля, за ним следует тип
значения, которое будет храниться в этом поле.
*/
package main

import (
	"fmt"
)

func main() {
	var myStruct struct {
		number float64
		word   string
		tgl    bool
	}
	fmt.Printf("%#v\n", myStruct) // Printf с глаголом %#v, она выведет значение из myStruct в виде литерала структуры
	// => struct { number float64; word string; tgl bool }{number:0, word:"", tgl:false}
	myStruct.number = 3.14       // присваиваем полю "number" значение 3.14
	fmt.Println(myStruct.number) // => 3.14
	myStruct.word = "pie"
	myStruct.tgl = true
}

/*
Определения типов позволяют создавать собственные типы. Они создают
новый определяемый тип на основе базового типа.
*/
type part struct { // struct - базовый тип для типа "part"
	description string
	count int
}

type car struct { // базовым типом для car таж же является структура с полями "name" и "speed"
	name string
	speed float64
}

func main() {
	var porsche car
	porsche.name = "Porsche 911"
	porsche.speed = 344
	fmt.Println("Name:", porsche.name)
	fmt.Println("Speed:", porsche.speed)
	
	var bolts part
	bolts.description = "Hex bolts"
	bolts.count = 24
	fmt.Println("Discription:", bolts.description)
	fmt.Println("Count:", bolts.count)
}