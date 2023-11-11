package flowcontrol

import "fmt"

func Sum(){
	sum:=1
	for sum < 100{
		sum += sum
		fmt.Println(sum)

	}
}
