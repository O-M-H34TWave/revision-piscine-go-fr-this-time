package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	result := 0
	if len(os.Args) != 2 {
		return
	}
	input := os.Args[1]
	i, _ := strconv.Atoi(input)
	for x := 1; x <= i; x++ {
		if i%x == 0 {
			result += i
		}
	}
	fmt.Println(result)
}
