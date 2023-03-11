package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i = ", i)
	}

	for j := 0; j < 11; j++ {
		if j == 5 {
			str := "CAШAPBO"
			for k, ch := range str {
				fmt.Printf("character %U '%c' starts at byte position %d\n", ch, ch, k*2)
			}
			continue
		}
		fmt.Println("Nilai j = ", j)
	}

	
}