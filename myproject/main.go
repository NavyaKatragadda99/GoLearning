package main

import (
	"example/flowcontrol"
	"fmt"
)
func main() {
    a := 87
    day := 3
    fmt.Println("Hello, World!")
    flowcontrol.Sum()
    flowcontrol.CheckEvenOrOdd(a)
    flowcontrol.DayOfWeek(day)
    fmt.Println(&a) //Pointer
    fmt.Println(flowcontrol.CreatePerson())
    flowcontrol.ReturnFirstName(flowcontrol.CreatePerson())
    flowcontrol.PrintArray()

}