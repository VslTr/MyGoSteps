Методы - функции связанные со значениями конкретного типа
у метода перед име­нем функции добавляется один дополнительный
параметр, называемый параметром получателя
/*
func (m MyType) sayHi() {
	fmt.Println("Hi from", m)
	}

(m MyType)
"m" - имя параметра получателя
"MyType" - тип параметра получателя */ 
Вызов метода, который вы определили, состоит из значения, для которого вызы­
вается метод, точки, имени вызываемого метода и пары круглых скобок. Значение,
для которого вызывается метод, называется получателем метода.
/*
value := MyType("a MyType value")
value.sayHi() // value - получатель метода, sayHi - Имя метода
*/
Метод может иметь так же обычные параметры и возвращать значения
/*
type MyType string

func (m MyType) MethodWithParameters(number int, flag bool) int {
	fmt.Printf("m: %s\n", m)
	fmt.Printf("number: %v\n", number)
	fmt.Printf("Flag: %v\n", flag)
	return len(m)
}

func main() {
	value := MyType("123456789")
	fmt.Printf("Len value: %v\n", value.MethodWithParameters(4, true))
}
// => m: 123456789
// => number: 4
// => Flag: true
// => Len value: 9
*/