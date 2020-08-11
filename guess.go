// guess - игра, в которой игрок должен угадать случайное число.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix() // получаем текущую дату и время в формате целого числа
	// Чтобы получать разные случайные числа, необходимо передать значение функции rand.Seed
	// Тем самым вы «инициализируйте» генератор случайных чисел, то есть предоставляете
	// значение, которое будет использоваться для генерирования других случайных чисел.
	// fmt.Println(seconds)
	/*
		Функция rand.Seed ожидает получить целое число, поэтому передать ей значение
		Time напрямую не удастся. Вместо этого для Time следует вызвать метод Unix, кото-
		рый преобразует его в целое число. (А конкретно значение будет преобразовано в
		формат времени Unix — целое количество секунд, прошедших с 1 января 1970 года.
	*/
	rand.Seed(seconds)           // функция генератора случайных чисел
	target := rand.Intn(100) + 1 // теперь генерируеммые числа будут разными при каждом зпуске

	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	reader := bufio.NewReader(os.Stdin) // создаем bufio.Reader для чтения ввода

	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")

		fmt.Print("Make a guess: ")           // запрос числа
		input, err := reader.ReadString('\n') // читаем двнные введенные до нажатия Enter
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)  // удаляем символ новой строки
		guess, err := strconv.Atoi(input) // преобразуем в целое число
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Oops. Your guess was LOW.") // введенное число меньше загаданного
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH.") // введенное число больше загаданного
		} else { // в противном случае введенное число правильное
			success = true                           // предотвращаем вывод сообщения о проигрыше
			fmt.Println("Good job! You guessed it!") // сообщаем что введенное число верное
			break                                    // выходим из цыкла
		}
	}

	if !success {
		fmt.Println("Sorry, you didn't guess my number. It was:", target) // если переменная "success" равна false, сообщаем игроку загаданное число
	}
	/*
		После цикла добавляется блок if для вывода сообщения о проигрыше. Но блок if выполняется только
		в том случае, если условие равно true, а мы хотим, чтобы сообщение о проигрыше выводилось только
		для переменной success, равной false. По этой причине мы добавляем оператор логического отри-
		цания (!). Как было показано ранее, этот оператор преобразует true в false, а false в true.
		В результате сообщение о проигрыше будет выводиться, если переменная success равна false, и не
		будет, если переменная success равна true.
	*/

}
