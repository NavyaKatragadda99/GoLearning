package SliceAndMap

import "fmt"

func CreateAndManipulateSlice() []int {
	slice := []int{1,2,3,4,5}
	slice = append(slice, 6)

	for i:=0; i<len(slice); i++ {
		slice[i] *=2
	}
	return slice
}

func PrintSlice(slice []int){
	fmt.Println(slice)
}