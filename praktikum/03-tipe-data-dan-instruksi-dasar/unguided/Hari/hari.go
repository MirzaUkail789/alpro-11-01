package main

import "fmt"

func main() {
	var hari int

	fmt.Scan(&hari)

	tahun := hari / 360
	hari = hari % 360

	bulan := hari / 30
	hari = hari % 30

	minggu := hari / 7
	sisa := hari % 7

	fmt.Println(tahun)
	fmt.Println(bulan)
	fmt.Println(minggu)
	fmt.Println(sisa)
}