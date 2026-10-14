package main
import "fmt"

func main() {
	var name string

	name = "Mirza Ukail Falah Triyarso"
	fmt.Println("Name : ", name)

	var lastName string = "Triyarso"
	fmt.Println("Last Name : ", lastName)

	middleName := "Ukail Falah"
	fmt.Println("Middle Name : ", middleName)

	var (
		fulName   string = "Mirza Ukail Falah Triyarso"
		firstName string = "Mirza"
	)
	fmt.Println(fulName)
	fmt.Println(firstName)
}