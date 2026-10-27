package main

import "fmt"

func main() {
	var n, sum int

	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	i := 1
	sum = 0

	for i <= n {
		sum += i
		i++
	}

	fmt.Println(sum)
}