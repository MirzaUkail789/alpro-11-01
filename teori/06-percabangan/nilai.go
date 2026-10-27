package main

import "fmt"

func main() {
	var nilai int
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nilai)

	var grade string

	if nilai < 50 {
		grade = "E"
	} else if nilai < 70 {
		grade = "c"
	} else if nilai < 85 {
		grade = "B"
	} else if nilai < 100 {
		grade = "A"
	}
	fmt.Println("Grade Anda:", grade)
}