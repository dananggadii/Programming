package main

import "fmt"

func main() {
	// Membuat array cara pertama
	var animals [3]string

	animals[0] = "Goat"
	animals[1] = "Lion"
	animals[2] = "Zebra"

	fmt.Println(animals[0])
	fmt.Println(animals[1])
	fmt.Println(animals[2])

	// Membuat array cara kedua (jika value array kurang, maka untuk melengkapi default nya di set 0)
	var country = [3]string{
		"German",
		"Indonesia",
		"Japan",
	}

	// print seluruh value pada array
	fmt.Println(country)
	// print value berdasarkan index
	fmt.Println(country[0])
	fmt.Println(country[1])
	fmt.Println(country[2])
	// Mengetahui panjang array
	fmt.Println(len(country))
	// Mengubah value pada array
	country[2] = "Singapore"
	fmt.Println(country[2])
}
