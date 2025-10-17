package main

import "fmt"

type Profile struct {
    Name    string
    City    string
    Campus  string
}

func (p *Profile) Show() {
    // Должен безопасно печатать «пустой профиль Интукод в Грозном», если p == nil
    if p == nil {
		fmt.Println("пустой профиль Интукод в Грозном")
	} else {
		fmt.Printf("Профиль: %s · %s (%s)\n", p.Name, p.City, p.Campus)
	}
}

func main() {
    var a *Profile = nil
    a.Show() // не должно паниковать

    b := &Profile{Name: "Кудузов Ахмад", City: "Грозный", Campus: "Школа программирования"}
    b.Show()
}