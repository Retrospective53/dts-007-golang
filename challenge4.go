package main

import (
	"fmt"
	"os"
	"strconv"
)

type student struct{
	Nama string
	Alamat string
	Pekerjaan string
	Alasan string
	Absen int
}

func main() {
	students := []student{
		{Nama: "Golang WOngsawat", Alamat: "thailand", Pekerjaan: "pro boxer", Alasan: "muh boxing", Absen: 1},
		{Nama: "Budiz", Alamat: "azkaban", Pekerjaan: "assassin", Alasan: "no diea llol", Absen: 5},
    {Nama: "Idubz", Alamat: "azkaban", Pekerjaan: "pro gamer", Alasan: "no diea llol", Absen: 2},
    {Nama: "El Donte", Alamat: "azkaban", Pekerjaan: "el diablo honter", Alasan: "no diea llol", Absen: 3},
    {Nama: "Dante", Alamat: "azkaban", Pekerjaan: "devil hunter", Alasan: "no diea llol", Absen: 4},
	}

	// fmt.Printf("the type is %T\n", students[0])

	absen, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid input argument")
	}

	var selectedStudent *student
	for i := range students {
		if students[i].Absen == absen {
			selectedStudent = &students[i]
			break
		}
	}


	// selectedStudent, ok := students[absen]
	// if ok {
	// 	fmt.Println("Nama: ", selectedStudent.Nama)
	// 	fmt.Println("Pekerjaan: ", selectedStudent.Pekerjaan)
	// 	fmt.Println("Alamat: ", selectedStudent.Alamat)
	// 	fmt.Println("Alasan: ", selectedStudent.Alasan)
	// } else {
	// 	fmt.Println("No student found with the given absen value")
	// }

	if selectedStudent != nil {
		fmt.Println("Nama: ", selectedStudent.Nama)
		fmt.Println("Pekerjaan: ", selectedStudent.Pekerjaan)
		fmt.Println("Alamat: ", selectedStudent.Alamat)
		fmt.Println("Alasan: ", selectedStudent.Alasan)
	} else {
		fmt.Println("No student found with the given absen value")
	}

}
