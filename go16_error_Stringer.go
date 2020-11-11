Встроенные в Go интерфейсы

Error
// Значение ошибки — это любое значение с методом Error, который возвращает строку
/*
Тип error является «предварительно объявленным идентификатором», как int или string.
И как все остальные предварительно объявленные идентификаторы, он не является частью
никакого пакета. Он является частью «универсального блока»; это означает, что
он доступен везде, независимо от того, какой пакет является текущим.
*/

// err := fmt.Errorf("a height of %0.2f is invalid", -2.33333)
// fmt.Println(err.Error()) // выведет => a height of -2.33 is invalid
// fmt.Println(err) // тоже выведет => a height of -2.33 is invalid

/*
Объявление типа error как интерфейса означает, что если тип содержит метод
Error, который возвращает string, то он поддерживает интерфейс error и может
использоваться в качестве значения ошибки. А это означает, что вы можете опре-
делять собственные типы и использовать их везде, где требуется значение ошибки!
---------------------------------------------------------------------------------------
Если кроме значения ошибки требуется еще и отслеживать более подробную информацию о ней
в формате, отличном от строки, вы можете создать свой собственный тип, поддерживающий
интерфейс error, и сохранить в нем нужную информацию.
Допустим, вы пишете программу, которая следит за некоторым устройством и предотвращает
его перегрев. Ниже приведен тип OverheatError, который помогает в решении этой задачи.
Он содержит метод Error и поэтому поддерживает error. Но при этом в качестве базового
типа используется float64, что позволяет нам отслеживать степень перегрева.
*/
package main

import (
	"fmt"
	"log"
)

type OverheatError float64
func (o OverheatError) Error() string {
	return fmt.Sprintf("Overheating by %0.2f degrees!", o)
}

func chekTemperature (actual float64, safe float64) error {
	excess := actual - safe 
	if excess > 0 {
		return OverheatError(excess)
	}
	return nil
}

func main() {
	var err error = chekTemperature(121.379, 100.0)
	if err != nil {
		log.Fatal(err)
	}
}
// Out => Overheating by 21.38 degrees!
---------------------------------------------------------------------------
интерфейс Stringer
При помощи этого интерфейса любой тип может отображаться при выводе.
Тоесть любой тип легко подготовить для поддержки Stringer,
достаточно определить метод String(), который возвращает строку.
// Определение интерфейса выглядит так
type Stringer interface {
	String() string
}
/*
Перейдем к примеру более серьезного использования типа интерфейса.
Сделаем так, чтобы наши типы Gallons, Liters и Milliliters поддерживали Stringer.
Мы переместим свой код форматирования этих значенийв методы String, связанные с каждым типом.
Вместо Printf метод будет вызывать функцию Sprintf и возвращать полученное значение.
 */
package main

import "fmt"

type  Gallons float64
func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gal", g)
}

type Liters float64
func (l Liters) String() string {
	return fmt.Sprintf("%0.2f L", l)
}

type Milliliters float64
func (m Milliliters) String() string {
	return fmt.Sprintf("%0.2f mL", m)
}

func main()  {
	fmt.Println(Gallons(12.09248342))
	fmt.Println(Liters(12.09248342))
	fmt.Println(Milliliters(12.09248342))
}
//output:
//		12.09 gal
//		12.09 L
//		12.09 mL

