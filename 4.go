package main

import (
	"fmt"
)

func cetakGambar() {
	str := []string{"P", "R", "O", "G", "R", "A", "M", "M", "E", "R"}
	for i := 0; i < len(str); i++ {
		for j := 0; j < len(str); j++ {
			if i == j || j == len(str)-i-1 {
				fmt.Print(str[i])
			} else {
				fmt.Print("=")
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	cetakGambar()
}
