package main

import "fmt"

func main() {
	lamp_status := "kunings"

	if lamp_status == "merah" {
		fmt.Println("Berhenti")
	} else if lamp_status == "kuning" {
		fmt.Println("Hati-hati")
	} else if lamp_status == "hijau" {
		fmt.Println("Jalan")
	} else  {
		fmt.Println("perhatian tidak terdaftar disini")
	}
}