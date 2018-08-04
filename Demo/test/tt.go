package main

import "fmt"

func sli(s []int) string {
	isBool := "A"
	for item := range s {
		fmt.Println(item)
		if item == 3 {
			isBool = "B"
			return "c"
		}
	}
	return isBool
}

func main() {
	test1 := []int{1, 2, 3, 4, 5}
	test2 := sli(test1)
	fmt.Println(test2)
}
