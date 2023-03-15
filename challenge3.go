package main

import "fmt"

func main() {
	str := "selamat malam"
	count := make(map[rune]int)
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}
	for _, char := range str {
		count[char]++
	}
	var countStr string
	for char, c := range count {
			countStr += fmt.Sprintf("%c:%d ", char, c)
	}
	countStr = "[ " + countStr + "]"
	fmt.Printf("%s", "map ")
	fmt.Println(countStr)

}
