package main

import "fmt"

func main() {
	var a, b int
	//membaca input
	fmt.Scan(&a)
	fmt.Scan(&b)

	//menukar nilai a dan b
	a, b = b, a
	//menampilkan output
	fmt.Println(a)
	fmt.Println(b)
}