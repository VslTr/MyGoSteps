// Форматирование вывода функциями Printf и Sprintf

/*
	функция Printf (сокращение от «print, with formatting», то есть «вывод с форматированием»).
	Функция получает строку и вставляет в нее одно или несколько значений, отформа-
	тированных заданным способом. После этого функция выводит полученную строку.
	fmt.Printf("About one-third: %0.2f\n", 1.0/3.0) // Output: About one-third: 0.33
	%0.2f - %(начало специффикатора форматирования), 0 - минимальная ширина всего числа, 2 - ширина дробной части, f - тип глагола форматирования
	Минимальная ширина всего числа включает дробную часть и точку разделитель.
	Если она указана, то более короткие числа будут дополняться пробелами в начале до достижения указанной ширины. Если она не указана (%.2f\n), то пробелы не добавляются.
	(%5.2f\n, 0.335656 => Output: пробел0.33 (пять символов спробелом))
	------------------------------------------------------------------------------------------
	Функция Sprintf (также из пакета fmt) в целом похожа на Printf, но она возвращает
	отформатированную строку, а не выводит ее.
	resultString := fmt.Sprintf("About one-third: %0.2f\n", 1.0/3.0)
	fmt.Printf(resultString) // Output: About one-third: 0.33
	------------------------------------------------------------------------------------------
	знаки процента (%) рассматриваются как начало глагола форматирования — части
	строки, которая будет заменена значением в определенном формате.
	fmt.Printf("The %s cost %d cents each.\n", "gumballs", 23)
	здесь символы со знаком процента будут заменены следующими по очереди значениями
	Output: The gumballs cost 23 cents each => %s заменяет ("gumballs"), а %d заменяет (23)
	------------------------------------------------------------------------------------------
	Буква, следующая за знаком процента, определяет тип глагола (показанны некоторые типы):
	%f - Число с плавающей точкой
	%d - Десятичное целое число
	%s - Строка
	%t - Логическое значение (true или false)
	%v - Произвольное значение (подходящий формат выбирается на основании типа передаваемого значения)
	%#v - Произвольное значение, отформатированное в том виде, в котором оно отображается в коде Go
	%T - Тип переданного значения (int, string и т. п.)
	%% - Знак процента (литерал)

	глагол формата %#v выводит все значения как есть (пустую строку, символ табуляции, символ новой строки и т.д)
	fmt.Printf("%#v %#v %#v", "", "\t", "\n") // Output: "" "\t" "\n"
*/

package main

import "fmt"

func paintNeeded(width float64, height float64) (float64, error) { // Если функция возвращает значение обязательно после переменных указавается тип возвращаемого значение (float64, error)
	if width < 0 { // если переданное число отрицательное (меньше 0), возвращаем 0 исообщение о ошибке
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 { // если переданное число отрицательное (меньше 0), возвращаем 0 исообщение о ошибке
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}
	area := width * height
	return area / 10.0, nil
	// возвращать надо значение такого же типа что и обявляли
}

func main() {
	amount, err := paintNeeded(4.2, -3.0)
	if err != nil {
		fmt.Println(err) // если в место принта вызвать log.Fatal(err), то так же выведится ошибка но дальнейшее выполнение программы прервется
	} else {
		fmt.Printf("%.2f liters needed\n", amount)
	}
	total := amount
	amount, err = paintNeeded(5.2, 3.55)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%.2f liters needed\n", amount)
	}

	total += amount
	fmt.Printf("Total: %.3f liters\n", total)

}

// Output: a height of -3.00 is invalid
// Output: 1.85 liters needed
// Output: Total: 1.846 liters