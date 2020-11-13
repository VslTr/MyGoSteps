package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)
// OpenFile() - Открывает файл и возвращает указатель на него, а так же значение обнаруженной ошибки
func  OpenFile(fileName string) (*os.File, error)  {
	fmt.Println("Opening", fileName)
	return os.Open(fileName)
}
// CloseFile() - Закрывает файл
func CloseFile(file *os.File)  {
	fmt.Println("Close file")
	file.Close()
}

func  GetFloats(fileName string) ([]float64, error)  {
	var numbers []float64 // слайс для хранения прочитаных чисел
	file, err := OpenFile(fileName) // Вместо прямого вызова os.Open, вызывается ф-я OpenFile()
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// в переменную "number" помещаем отсканированную и сконвертированную в float64 строку.
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number) // добавляем число из переменной "namber" в слайс
	}
	CloseFile(file) // Вместо прямого вызова file.Close, вызывается CloseFil()
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil // возвращаем слайс с отсканированными и сконвертированными значениями, а также nil если нет ошибки
}

func main()  {
	// run readfile.go data.txt (data.txt - первый аргумент командной строки)
	numbers, err := GetFloats(os.Args[1])// "Args[1]" первый аргумент командной строки используется как имя файла
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers { // берем из слайса значения
		fmt.Println(number)
		sum += number // суммируем
	}
	fmt.Printf("Sum: %0.2f\n", sum)
}

// Запустите программу командой go run readfile.go data.txt


//----------------------------------------------------------------------------------------------
//func main() {
	//// Открываем файл для чтения
	//file, err := os.Open("data.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//// Для файла создается значение "Scanner" и присваивается переменной "scanner"
	//scanner := bufio.NewScanner(file)
	//
	//// Цикл выполняется пока не будет достигнут конец файла
	//// и "scanner.Scan" не вернет "false"
	//for scanner.Scan() { // Читаем строку из файла
	//	fmt.Println(scanner.Text()) // Выводим строку
	//}
	//// Если файл останется открытым, он расходует ресурсы операционной системы,
    //// поэтому файлы всегда должны закрываться после завершения работы с ними
	//err = file.Close() // Закрываем файл
	//if err != nil { // Если при закрытии файла произошла ошибка,
	//	log.Fatal(err) // сообщить о ней и завершить работу
	//}
	//
	//
	//if scanner.Err() != nil { // Если при сканировании файла произошла ошибка,
	//	log.Fatal(scanner.Err()) // сообщить о ней и завершить работу
	//}
	//}