package main

import "fmt"

func main() {
	var celsius, reamur float64

	fmt.Scan(&celsius)
	reamur = 4.0 / 5.0 * celsius
	fmt.Println(reamur)
}