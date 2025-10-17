package main

import "fmt"

type Passport struct {
	FirstName string
	LastName string
	Active bool
}

func (pass Passport) FullName() string {
	return pass.FirstName + pass.LastName
}

func (pass *Passport) Deactivate() {
	pass.Active = false
}

func (pass Passport) IsActive() {
	fmt.Println(pass.Active)
}

func main() {
	newPass := Passport{
		FirstName: "Bob",
		LastName: "Smit",
		Active: true,
	}
	fmt.Println(newPass.FullName())
	newPass.Deactivate()
	newPass.IsActive()

}