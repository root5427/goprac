package main

import (
	"fmt"
	"goprac/begin"
	"goprac/integer"
)

func main() {
	var catNo int = 0
	fmt.Println("This is go practice!")
	fmt.Println("Choose the category: ")
	fmt.Println("1. Begin")
	fmt.Println("2. Integer")
	fmt.Println("3. Boolean")
	fmt.Println("5. If")
	fmt.Println("6. Case")
	fmt.Println("7. For")
	fmt.Print("Enter the category number: ")
	fmt.Scan(&catNo)
	if catNo > 0 {
		switch catNo {
		case 1:
			begin.CallBegin()
		case 2:
			integer.CallInteger()
		}
	}
}
