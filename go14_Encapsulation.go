Инкапсуляция
/*
Практика сокрытия данных в одной части программы от кода в другой части называется
инкапсуляцией.
Многие другие языки программирования инкапсулируют данные в классах. (На концептуаль-
ном уровне классы похожи на типы Go, но не идентичны им.) В Go данные инкасулируются в
пакетах с применением неэкспортируемых переменных, полей структур, функций или методов.
*/
Для доступа к инкапсулированным (скрытых) в пакетах данным используют get- и set-методы
(геттеры и сеттеры)
/*
Сообщество Go выработало соглашения, по которым префикс
Get не указывается в именах get-методов. Включение префикса
только запутает ваших коллег-разработчиков!
Для set-методов в Go используется префикс Set, как и во мно-
гих других языках, потому что позволяет отличить set-методы от
get-методов для того же поля.
*/
ПРИМЕР
/*
// Создадим отдельный файл "date.go" в пакете "calendar"

package calendar

import "errors"
// тип Date с недоступными из вне полями (имя полей начинается с маленьких букв)
type Date struct {
	year  int
	month int
	day   int
}
// Функции сеттеры позволяющие нам задавать параметры полей типа Date из других пакетов
// Имя функций начинается с большой буквы
func (d *Date) SetYear(year int) error { 
	if year > 1900 && year <= 2020 {
		d.year = year
		return nil
	} else {
		return errors.New("Invalid year")
	}
}
func (d *Date) SetMonth(month int) error {
	if month > 0 && month < 13 {
		d.month = month
		return nil
	} else {
		return errors.New("Invalid month")
	}
}
func (d *Date) SetDay(day int) error {
	if day > 0 && day < 32 {
		d.day = day
		return nil
	} else {
		return errors.New("Invalid day")
	}
}

// Функции геттеры позволяющие получить данные полей типа Date из вне
func (d *Date) Year() int {
	return d.year
}
func (d *Date) Month() int {
	return d.month
}
func (d *Date) Day() int {
	return d.day
}
*/

package main

import (
	"fmt"
	"log"
	"github.com/headfirstgo/calendar"
)

func main() {
	date := calendar.Date{}
	err := date.SetYear(2020) // используем сеттер "SetYear" для передачи значения полю "year"
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetMonth(10)
	if err != nil {
	log.Fatal(err)
	}

	err = date.SetDay(1)
	if err != nil {
	log.Fatal(err)
	}

	fmt.Printf("%v.%v.%v\n", date.Day(), date.Month(), date.Year())
	// используем геттеры для получения значений полей
}
