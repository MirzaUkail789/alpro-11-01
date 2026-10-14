package main

import "fmt"

func main() {
	var n, jumlah int
	fmt.Scan(&n)

	for i := 1; i <= n; i += 2 {
		if i > 1 {
			fmt.Print(" ")
		}
		fmt.Print(i)
		jumlah += i
	}
	fmt.Println()
	fmt.Println(jumlah)
}