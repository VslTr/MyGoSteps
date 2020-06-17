// pass_fail сообщает, сдал ли пользователь экзамен.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	//input, _ := reader.ReadString('\n') // Присваиваум второму значению пустой идентификатор
	/*
		В большинстве языков программирования функции и методы
		могут возвращать только одно значение, но в Go функция
		может возвращать сколько угодно значений. Чаще всего мно-
		жественные возвращаемые значения в Go используются для
		возвращения дополнительного значения ошибки.
	*/
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err) // Если «ошибка» cообщить об ошибке и прервать
	}
	fmt.Println(input)
}
