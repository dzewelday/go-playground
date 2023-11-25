package main

import "fmt"

type Printer interface {
	Print()
}

type Claim struct {
	Id   int
	Type string
}

func (c Claim) Print() {
	fmt.Println(c.Id, c.Type)
}

type User struct {
	Id   int
	Name string
}

func (u User) Print() {
	fmt.Println(u.Id, u.Name)
}

func Tester() {
	p := Claim{Id: 1, Type: "Health"}
	p.Print()

	u := User{Id: 1, Name: "John"}
	u.Print()
}
