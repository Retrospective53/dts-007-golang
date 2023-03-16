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
		{Nama: "budi", Alamat: "azkaban", Pekerjaan: "assassin", Alasan: "no diea llol", Absen: 1},
    {Nama: "idubz", Alamat: "azkaban", Pekerjaan: "assassin", Alasan: "no diea llol", Absen: 2},
	}

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

	if selectedStudent != nil {
		fmt.Println(*selectedStudent)
	} else {
		fmt.Println("No student found with the given absen value")
	}

	// fmt.Println(budi)
	// fmt.Println(idubz)

	// lol := len(os.Args)
	// fmt.Println(lol)
}
