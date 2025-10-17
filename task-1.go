
package main

import "fmt"

type Student struct {
    Name string
    Age  int
}

// TODO: добавь метод Info() со значением-получателем (Student),
// который вернёт строку вида: "Студент: Имя (Возраст: N)"
// Подсказка: fmt.Sprintf поможет собрать строку.

func (s Student) Info() string {
	return fmt.Sprintf("Студент: %s (Возраст: %d)", s.Name, s.Age)
}

func main() {
    s := Student{Name: "Амина", Age: 18}
    // Ожидается печать строки с именем и возрастом
    fmt.Println(s.Info())
}