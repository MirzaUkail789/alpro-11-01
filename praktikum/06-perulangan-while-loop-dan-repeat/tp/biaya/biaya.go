package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan tahun ke-n: ")
	fmt.Scan(&n)

	// Biaya awal pada tahun ke-1
	var biaya float64 = 10000000

	i := 1
	for i < n {
		biaya = biaya + (biaya * 0.06)
		i++
	}

	fmt.Printf("%.0f\n", biaya)
}