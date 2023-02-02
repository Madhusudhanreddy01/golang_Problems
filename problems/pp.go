package main

import (
	"fmt"
)

func myFunction(w int, d int) int {
	return w * ^d
}

func main() {
	fmt.Println(myFunction(10, 20))
}
