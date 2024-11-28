package main

import "fmt"

func main() {
	for i := 9; i >= 2; i-- {
		for j := 8; j >= 1; j-- {
			for k := 7; k >= 0; k-- {
				if i != j && i != k && j != k {
					fmt.Printf("%d%d%d", i, j, k)
					if i != 2 || j != 1 || k != 0 {
						fmt.Printf(", ")
					}
				}
			}
		}
	}
}
