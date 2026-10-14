package main
import "fmt"

func main() {
	var posisi, posisi0, kecepatan, waktu float64
	fmt.Scan(&posisi0, &kecepatan, &waktu)
	posisi = posisi0 + kecepatan*waktu
	fmt.Println("Posisi atau Jarak adalah: ", posisi, "meter")
}
