package main
import "fmt"

// const (
// 	c1 = iota * 2
// 	c2
// 	c3
// )



func main() {
	i := 21
	j := true
	k := 123.456
	fmt.Printf("%d \n", i)
	fmt.Printf("%T \n", i)
	fmt.Printf("%s \n", "%")
	fmt.Printf("%t \n", j)
	fmt.Printf("\n")
	fmt.Printf("%b \n", 21)
	fmt.Printf("%c \n", 0x042F)
	fmt.Printf("%d \n", 21)
	fmt.Printf("%o \n", 21)
	fmt.Printf("%x \n", 15)
	fmt.Printf("%X \n", 15)
	fmt.Printf("%U \n", 'Я')
	fmt.Printf("%.3f \n", k)
	fmt.Printf("%f \n", k)
	fmt.Printf("%E \n", k)
}