package main

import "fmt"

func main() {
	var m int

	fmt.Scan(&m)

	var jumlahSertifikat int = 0

	for i := 0; i < m; i++ {
		var p, v, ga, fs, ci, tc float64

		fmt.Scan(&p, &v, &ga, &fs, &ci, &tc)

		rataRata := (p + v + ga + fs + ci + tc) / 6.0

		if rataRata >= 3.33 {
			fmt.Println("Ya")
			jumlahSertifikat++
		} else {
			fmt.Println("Tidak")
		}
	}

	fmt.Println(jumlahSertifikat)
}