package main

import "fmt"

func main() {
	var a, b int

	fmt.Scan(&a)
	fmt.Scan(&b)

	jumlah := a + b
	kurang := a - b
	kali := a * b
	bagi := a / b
	sisa :=   a % b

	fmt.Println(jumlah, kurang, kali, bagi, sisa)
}