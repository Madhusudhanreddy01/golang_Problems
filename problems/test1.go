package main

import "fmt"

func main() {
	var i int = 2
	fmt.Print(i)

	for i <= 5 {

		fmt.Print(" ", i*2)
		i++
	}

}
