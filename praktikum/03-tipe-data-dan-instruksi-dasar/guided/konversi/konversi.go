package main

import "fmt"

func main() {
	var celsius float64;

	fmt.Print("Masukkan suhu: ")
	fmt.Scan(&celsius)

	fmt.Println(celsius + 273)

}