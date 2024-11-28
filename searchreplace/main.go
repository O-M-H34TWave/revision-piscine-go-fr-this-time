package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	input := os.Args[1]
	old := os.Args[2]
	new := os.Args[3]
	res := ""
	for i := 0; i < len(input); i++ {
		if string(input[i]) == old {
			// ila index f iput equal to old, res zidlha 
			// new
			res += new
		} else {
			// else rj3lya gha string f index i w appendih l res
			res += string(input[i])
		}
	}
	fmt.Println(res)
}
