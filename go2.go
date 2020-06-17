package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	fmt.Println(math.Floor(16.75))
	// функция Floor число с плавающей точкой, округляет его до ближайшего меньшего целого и возвращает полученное число
	fmt.Println(strings.Title("head first go"))
	// функция Title преобразует первую букву каждого слова к верхнему регистру и возвращает полученную строку
	var txt string
	txt = "Gggg"
	nm := 42.34
	fmt.Println("nm is: ", reflect.TypeOf(nm), "\ntxt is: ", reflect.TypeOf(txt))
	// TypeOf из пакета reflect  покажет тип переменной
}
