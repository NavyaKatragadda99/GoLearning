package flowcontrol

import "fmt"

func CheckEvenOrOdd(a int){
	if a%2 == 0 {
		fmt.Println("a is even")
	} else {
		fmt.Println("a is odd")
	}
}
