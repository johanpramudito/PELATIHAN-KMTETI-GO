package main

import (
	"fmt"
)

func main() {
	var celsius float64
	var pilihan int

	// Input suhu dalam Celsius
	fmt.Print("Masukkan suhu dalam Celsius: ")
	fmt.Scan(&celsius)

	// Menu konversi
	fmt.Println("Pilih konversi:")
	fmt.Println("1. Fahrenheit")
	fmt.Println("2. Reamur")
	fmt.Println("3. Kelvin")
	fmt.Print("Masukkan pilihan (1/2/3): ")
	fmt.Scan(&pilihan)

	// Switch case untuk memilih konversi
	switch pilihan {
	case 1:
		// Celsius ke Fahrenheit
		fahrenheit := (celsius * 9 / 5) + 32
		fmt.Printf("%.2f Celsius = %.2f Fahrenheit\n", celsius, fahrenheit)
	case 2:
		// Celsius ke Reamur
		reamur := celsius * 4 / 5
		fmt.Printf("%.2f Celsius = %.2f Reamur\n", celsius, reamur)
	case 3:
		// Celsius ke Kelvin
		kelvin := celsius + 273
		fmt.Printf("%.2f Celsius = %.2f Kelvin\n", celsius, kelvin)
	default:
		fmt.Println("Pilihan tidak valid!")
	}
}
