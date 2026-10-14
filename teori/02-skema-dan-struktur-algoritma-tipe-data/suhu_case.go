package main
import "fmt"

func main() {
	var umur int8
	var suhu float32

	suhu = 36.3
	umur = 10

	fmt.Println("Umur : ", umur)
	fmt.Println("Suhu Tubuh : ", suhu)
	fmt.Println("Alamat memori dari var suhu ", &suhu) 
	fmt.Println("Alamat memori dari var umur ", &umur) 
}