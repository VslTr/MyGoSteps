package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Открываем файл для чтения
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Для файла создается значение "Scanner" и присваивается переменной "scanner"
	scanner := bufio.NewScanner(file)
	
	// Цикл выполняется пока не будет достигнут конец файла 
	// и "scanner.Scan" не вернет "false"
	for scanner.Scan() { // Читаем строку из файла
		fmt.Println(scanner.Text()) // Выводим строку
	}
	// Если файл останется открытым, он расходует ресурсы операционной системы,
    // поэтому файлы всегда должны закрываться после завершения работы с ними
	err = file.Close() // Закрываем файл
	if err != nil { // Если при закрытии файла произошла ошибка,
		log.Fatal(err) // сообщить о ней и завершить работу
	}
	
	
	if scanner.Err() != nil { // Если при сканировании файла произошла ошибка,
		log.Fatal(scanner.Err()) // сообщить о ней и завершить работу
	}
}