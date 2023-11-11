package flowcontrol

import "fmt"

type Person struct {
	FirstName string
	LastName string
	Age int
}

func CreatePerson() Person {
	p := Person{FirstName: "Maeve", LastName: "Weilly", Age: 32}
	return p
}

func ReturnFirstName(p Person){
	fmt.Println(p.FirstName)
}