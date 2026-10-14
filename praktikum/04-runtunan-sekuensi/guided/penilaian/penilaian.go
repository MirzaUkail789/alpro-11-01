package main

import "bufio"
import "fmt"
import "os"

func main() {
	var (
		nama, grade    string
		nilai, pilihan int
	)

	fmt.Println("============= Menu =============")
	fmt.Println("1. Sistem penilaian")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih opsi: ")
	fmt.Scanln(&pilihan)

	if pilihan == 1 {
		fmt.Print("Masukkan Nama:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		nama = scanner.Text()

		fmt.Print("Masukkan Nilai:")
		fmt.Scan(&nilai)

		if nilai >= 90 && nilai <= 100 {
			grade = "A"
		} else if nilai >= 80 && nilai < 90 {
			grade = "B"
		} else if nilai >= 70 && nilai < 80 {
			grade = "C"
		} else if nilai >= 60 && nilai < 70 {
			grade = "D"
		} else {
			grade = "E"
		}

		fmt.Println(nama)
		fmt.Println(grade)
	} else if pilihan == 0 {
		return
	} else {
		fmt.Println("Pilihan tidak valid")
	}
}