// pass_fail сообщает, сдал ли пользователь экзамен.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv" // для ParseFloat
	"strings" // для TrimSpase
)

func main() {
	fmt.Print("Enter a grade: ") // запрашиваем у пользователя значение
	var status string
	reader := bufio.NewReader(os.Stdin) // создаем bufio.Reader для чтения ввода с клавиатуры
	//input, _ := reader.ReadString('\n') // Присваиваум второму значению пустой идентификатор
	/*
		В большинстве языков программирования функции и методы
		могут возвращать только одно значение, но в Go функция
		может возвращать сколько угодно значений. Чаще всего мно-
		жественные возвращаемые значения в Go используются для
		возвращения дополнительного значения ошибки.
	*/
	input, err := reader.ReadString('\n') // читаем данные вводимые пользователем до нажатия Enter
	if err != nil {
		log.Fatal(err) // Если «ошибка» cообщить об ошибке и прервать
	}

	input = strings.TrimSpace(input)            // удаляем символы новой строки ('\n') и пробелы по краям (если есть) из введенных данных
	grade, err := strconv.ParseFloat(input, 64) // конвертируем из строки в float64
	if err != nil {
		log.Fatal(err)
	}

	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}

	fmt.Println("A grade of", grade, "is", status)
}

// Output:
// Enter a grade: 60
// A grade of 60 is passing

//-----------------------------------------------------------------------------------------------------------------------------------
/*
// Узнаем размер файла (в байтах)
import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileInfo, err := os.Stat("my.pdf")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfo.Size())
	// Output:  38631014 (byte)
}
*/
