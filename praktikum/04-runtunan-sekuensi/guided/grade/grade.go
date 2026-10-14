package main
import "bufio"
import "fmt"
import "os"

func main() {
	var (
		nilai float64
		nama,grade string
	)
	// bufio scanner untuk input
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("Masukkan nama: ")
	scanner.Scan()
	nama = scanner.Text()

	fmt.Println("Masukkan nilai: ")
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
}