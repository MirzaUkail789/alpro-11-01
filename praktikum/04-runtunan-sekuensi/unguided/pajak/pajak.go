package main

import "fmt"

func main() {
	var (
		penghasilan, pajak float64
	)

	// Persents
	p5 := 5.0 / 100.0
	p10 := 10.0 / 100.0
	p15 := 15.0 / 100.0
	p20 := 20.0 / 100.0

	fmt.Print("Masukkan penghasilan (juta): ")
	fmt.Scan(&penghasilan)

	switch {
	case penghasilan <= 50:
		pajak = penghasilan * p5
		fmt.Printf("Total pajak: %.1f juta\n", pajak)

	case penghasilan > 50 && penghasilan <= 100:
		pajak = (50 * p5) + ((penghasilan - 50) * p10)
		fmt.Printf("Total pajak: %.1f juta\n", pajak)

	case penghasilan > 100 && penghasilan <= 200:
		pajak = (50 * p5) + (50 * p10) + ((penghasilan - 100) * p15)
		fmt.Printf("Total pajak: %.1f juta\n", pajak)

	case penghasilan > 200:
		pajak = (50 * p5) + (50 * p10) + (100 * p15) + ((penghasilan - 200) * p20)
		fmt.Printf("Total pajak: %.1f juta\n", pajak)
	}
}