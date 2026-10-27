package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var nama, min1Nama, min2Nama string
	var nilai, min1Nilai, min2Nilai int

	i := 1
	for i <= n {
		fmt.Scan(&nama, &nilai)

		if i == 1 {
			min1Nilai = nilai
			min1Nama = nama
			min2Nilai = nilai
			min2Nama = nama
		} else {
			if nilai < min1Nilai {
				min2Nilai = min1Nilai
				min2Nama = min1Nama
				min1Nilai = nilai
				min1Nama = nama
			} else if min1Nilai == min2Nilai || nilai < min2Nilai {
				if min1Nilai == min2Nilai && nilai > min1Nilai {
					min2Nilai = nilai
					min2Nama = nama
				} else if nilai != min1Nilai && (nilai < min2Nilai || min1Nilai == min2Nilai) {
					min2Nilai = nilai
					min2Nama = nama
				}
			}
		}
		i++
	}

	fmt.Println(min1Nama)
	fmt.Println(min2Nama)
}