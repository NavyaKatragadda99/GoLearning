package SliceAndMap

import "fmt"

func CreateAndManipulateMap() {
	m := make(map[string]int)
	m["key1"] = 10
	m["key2"] = 20
	m["key3"] = 30

	m["key2"] = 222

	v, ok :=m["key2"]
	fmt.Println("The value:", v, "present?", ok)
	i, j := m["key5"]
	fmt.Println("The value:", i, "present?", j)
}

