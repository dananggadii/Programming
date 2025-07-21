package main

import "fmt"

func main() {
	const (
		nilai_agama      = 80
		nilai_matematika = 90
		total_nilai      = nilai_agama + nilai_matematika
	)

	// Augmented assignments
	var nilai_a = 70
	nilai_a += 80

	// unary operator / shortcut augmented assignments
	var nilai_b = 40
	nilai_b++ //operasi berikut akan menambah 1 nilai dari variable (40 + 1)

	var nilai_c = 30
	nilai_c--

	fmt.Println("Total nilai danang pada ulangan hari ini adalah =", total_nilai)
	fmt.Println("Hasil dari nilai berikut adalah", nilai_a)
	fmt.Println("Hasil dari nilai berikut adalah", nilai_b)
	fmt.Println("Hasil dari nilai berikut adalah", nilai_c)
}
