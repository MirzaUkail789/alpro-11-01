package main

import "fmt"

func main() {
	// Inisialisasi data
	correctPassword := "1234"
	saldo := 10000000
	
	var password string
	var attempts int = 0
	var isLoggedIn bool = false

	// Proses login (maksimal 3 kesempatan)
	for attempts < 3 && !isLoggedIn {
		fmt.Print("Masukkan password ATM: ")
		fmt.Scan(&password)

		if password == correctPassword {
			isLoggedIn = true
			fmt.Println("Login berhasil!")
		} else {
			attempts++
			if attempts < 3 {
				fmt.Println("Password salah. Silakan coba lagi.")
			}
		}
	}

	// Jika salah 3 kali berturut-turut, ATM terblokir
	if !isLoggedIn {
		fmt.Println("ATM terblokir. Program selesai.")
		return
	}

	// Menu ATM (berulang selama pengguna belum memilih keluar)
	var pilihan int
	for {
		fmt.Println("Menu ATM:")
		fmt.Println("1. Lihat Saldo")
		fmt.Println("2. Tarik Uang")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			fmt.Printf("Saldo Anda: Rp %d\n", saldo)
		} else if pilihan == 2 {
			var tarik int
			fmt.Print("Masukkan jumlah uang yang ingin ditarik: Rp ")
			fmt.Scan(&tarik)

			if tarik <= saldo {
				saldo -= tarik
				fmt.Printf("Anda berhasil menarik uang sebesar Rp %d\n", tarik)
				fmt.Printf("Saldo tersisa: Rp %d\n", saldo)
			} else {
				fmt.Println("Saldo tidak mencukupi.")
			}
		} else if pilihan == 3 {
			fmt.Println("Keluar dari sistem. Terima kasih!")
			break
		} else {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}