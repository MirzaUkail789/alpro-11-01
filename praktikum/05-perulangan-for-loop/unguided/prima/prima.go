package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	var jumlahFaktor int = 0

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Printf("%d ", i)
			jumlahFaktor++
		}
	}
	fmt.Println()

	if jumlahFaktor == 2 {
		fmt.Println("Ya")
	} else {
		fmt.Println("Tidak")
	}
}