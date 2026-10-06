package main

import "fmt"

func main() {
	var x int
	fmt.Print("Masukkan nominal: ")
	fmt.Scan(&x)

	var sepuluhribu int = x / 10000
	x = x % 10000

	var limaribu int = x / 5000
	x = x % 5000

	var seribu int = x / 1000

	fmt.Println(sepuluhribu, "lembar")
	fmt.Println(limaribu, "lembar")
	fmt.Println(seribu, "lembar")
}