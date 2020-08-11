// УКАЗАТЕЛИ
/* Значения, представляющие адреса переменных, называются
указателями, потому что они указывают на область памяти,
в которой хранится переменная. */
// Оператор & (амперсанд) используется в Go для получения адреса переменной.
/*	amount := 6
	fmt.Println(amount) // Вывод значения переменной Output => 6
	fmt.Println(&amount) // Вывод адреса переменной Output => 0x1040a124	*/
// -------------------------------------------------------------------------------------------------
// С помощью функции reflect.TypeOf (import "reflect") можно вывести типы указателей
// var myInt int
// fmt.Println(reflect.TypeOf(&myInt))  // Output => *int
// -------------------------------------------------------------------------------------------------
// чтобы получить значение переменной по указателю, ставьте оператор * прямо перед указателем в коде
/*  myInt := 4
myIntPointer := &myInt
fmt.Println(myIntPointer) // Output => 0x1040a124
fmt.Println(*myIntPointer) // Output => 4  */
// -------------------------------------------------------------------------------------------------
// Оператор * может использоваться для обновления значения по указателю
/*	myInt := 4
	fmt.Println(myInt)	// Output => 4
	myIntPointer := &myInt	// Присваиваем переменной указатель
	*myIntPointer = 8	// Новое значение присваеваем переменной, на которую ссылается указатель.
	fmt.Println(*myIntPointer)	// Output => 8
	fmt.Println(myInt)	// Output => 8	*/
// --------------------------------------------------------------------------------------------------------
// Указатели также можно возвращать из функций, просто объявите возвращаемый тип функции как тип указателя.
/*
func createPointer() *float64 {	 // Обьявляем что функция возвращает указатель на float64
	var myFloat = 98.5
	return &myFloat
}

func main() {
	var myFloatPointer *float64 = createPointer() // Назначаем возвращенный указатель переменной.
	fmt.Println(*myFloatPointer)
}
*/

package main

import (
	"fmt"
)

func test(a *float64) (float64, error) { // получаем указатель "a" на переменную  "chVar"
	if *a < 0 { // Проверяем не пришло ли отрицательное число (пример проверки корректности ввода данных)
		return 0, fmt.Errorf("a %0.2f is invalid", a)
	}

	*a += 1.0                     // Меняем значение переменной "chVar" (chVar теперь имеет значение 11) на которое ссылается указатель "a"
	fmt.Println("Adress a = ", a) // Output => Adress a =  0xc0000140a0
	return *a - 3.0, nil          // Делаем еще некие операции со значением переменной "chVar" (11 - 3), возвращаем "8" и "nil".
}

func main() {
	chVar := 10.0
	//chVarPointer := &chVar
	fmt.Printf("Adress chVar: %v  chVar: %v", &chVar, chVar) // Output => Adress chVar: 0xc0000140a0  chVar: 10
	fmt.Println("\n")
	t, err := test(&chVar) // Передаем не значение переменной chVar, а указатель на нее. Получаем "8" (t = 8) и "nil" (err = nil).
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Adress t: %v  t: %v\n", t, &t)                               // Output => Adress t: 8  t: 0xc0000140b8
		fmt.Printf("Adress chVar: %v  chVar: %v\n", chVarPointer, *chVarPointer) // Output => Adress chVar: 0xc0000140a0  chVar: 11
	}
}
