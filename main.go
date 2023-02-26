package main

import "fmt"

func addTwoNum(x int, y int) (result int) {
	result = x + y
	return
}

func main() {
	sum := addTwoNum(10, 20)
	fmt.Println(sum)

	count := 5

	for i := 0; i < count; i++ {
		fmt.Printf("%v ", i)
	}

	fmt.Printf("\n")

	// fruits := [3]string{"apple", "orange", "banana"}
	// for idx, val := range fruits {
	// 	fmt.Printf("%v\t%v\n", idx, val)
	// }

	fruits := [3]string{"apple", "orange", "banana"}
	for _, val := range fruits {
		fmt.Printf("%v\n", val)
	}
}
