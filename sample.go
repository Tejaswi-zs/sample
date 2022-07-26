package main

import "fmt"

func main() {
	//m := make(map[string]int)
	//m1 := "MISSISSIPPI"
	//for _, v := range m1 {
	//	m[string(v)]++
	//}
	//fmt.Println(m)
	var my interface{}
	switch v := my.(type) {
	case int:
		fmt.Println(v)
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("mi mama")
	}

}
