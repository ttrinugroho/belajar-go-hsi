package main

import (
	"fmt"
	"strconv"
	"tugas-pertemuan-3/mahasiswa"
)

func sumUmur() (*int, func(int)) {
	umur := 0
	return &umur, func(i int) {
		umur += i
	}
}

func main() {
	umur, add := sumUmur()

	ali := mahasiswa.BuatMahasiswa("ali", 20, 98, 93, 77, 82, 83)
	fmt.Println(ali.Info())
	fmt.Println("Rata-Rata nilai: " + strconv.FormatFloat(ali.RataRata(), 'f', 2, 64))

	add(ali.GetUmur())
	fmt.Println("---")

	siti := mahasiswa.BuatMahasiswa("Siti", 22, 90, 99, 90, 99, 90, 99)
	fmt.Println(siti.Info())
	fmt.Println("Rata-Rata nilai: " + strconv.FormatFloat(siti.RataRata(), 'f', 2, 64))

	add(siti.GetUmur())
	fmt.Println("---")

	fmt.Println("Version Package:", mahasiswa.Version)
	fmt.Println("Nilai Maksimum:", strconv.Itoa(mahasiswa.GetMaxNilai()))
	fmt.Println("Total umur Mahasiswa:", strconv.Itoa(*umur))
}
