package main
import "fmt"

func main() {
	var (
		usia,gajiTahunan int
	)
	fmt.Print("Masukkan usia: ")
	fmt.Scan(&usia)

	fmt.Print("Masukkan gaji tahunan (juta): ")
	fmt.Scan(&gajiTahunan)

	switch {
		case usia < 18 && gajiTahunan == 0 :
			fmt.Println("Masih Sekolah")
		case (usia >= 18 && usia <= 25) && gajiTahunan >= 50 :
			fmt.Println("Muda Sukses")
		case (usia >= 18 && usia <= 25) && gajiTahunan < 50 :
			fmt.Println("Masih belajar hidup")
		case (usia >= 26 && usia <= 40) && gajiTahunan >= 100 :
			fmt.Println("Pekerja Mapan")
		case (usia >= 26 && usia <= 40) && gajiTahunan < 100 :
			fmt.Println("Perlu perbaikan karir")
		case usia > 40 && gajiTahunan >= 150 :
			fmt.Println("Profesional Berpengalaman")
		case usia > 40 && gajiTahunan < 150 :
			fmt.Println("Perlu Evaluasi Finansial")
	}
}