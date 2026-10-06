package main

import "fmt"

func main() {
	var nama string
	var skorMatematika, skorBahasaInggris  int

	//membaca input
	fmt.Scan(&nama)
	fmt.Scan(&skorMatematika)
	fmt.Scan(&skorBahasaInggris)

	//menghitung total & rata-rata
	total := skorMatematika + skorBahasaInggris
	rataRata := total / 2
	//menampilkan output
	fmt.Println(nama)
	fmt.Println(total)
	fmt.Println(rataRata)
}