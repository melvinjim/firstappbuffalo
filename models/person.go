package models

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

func (p Person) FullName() string {
	fullname := fmt.Sprintf("%v %v", p.FirstName, p.LastName)

	return fullname
}
