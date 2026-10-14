package main

import "fmt"

func main() {
	var mil float64

	fmt.Scan(&mil)

	kilometer := mil * 1.6

	fmt.Printf("%.1f\n", kilometer)
}