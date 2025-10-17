package main

import "fmt"

type Counter struct {
    value int
}

func (c *Counter) Inc() { // проблема была в том что передается копия значения а не ссылка в памяти
    c.value++
}

func (c * Counter) Reset() {
	c.value = 0
}

func main() {
    var day Counter
    day.Inc()
    day.Inc()
    fmt.Println("Оформлено заказов:", day.value) // ожидаем 2, получаем 0
}