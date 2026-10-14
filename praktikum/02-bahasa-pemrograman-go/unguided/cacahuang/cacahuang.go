package main

import "fmt"

func main() {
	var totalUang, sisaUang int
	var n10k, n5k, n1k int

	// Membaca input dari pengguna
	fmt.Scan(&totalUang)

	// Menghitung jumlah lembar menggunakan variabel sisa yang jelas
	n10k = totalUang / 10000
	sisaUang = totalUang % 10000

	n5k = sisaUang / 5000
	sisaUang = sisaUang % 5000

	n1k = sisaUang / 1000

	// Menampilkan hasil cetak dengan format satu baris (dipisah spasi)
	// Sesuai dengan format keluaran umum pada soal pemrograman dasar
	fmt.Printf("%d lembar 10000, %d lembar 5000, %d lembar 1000\n", n10k, n5k, n1k)
}