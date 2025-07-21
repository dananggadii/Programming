package main

import "fmt"

func main() {
	const (
		nilai_akhir   = 90
		nilai_absensi = 85
		nilai_lulus   = 80
	)

	var (
		cek_nilaiLulus   = nilai_akhir >= nilai_lulus
		cek_absensiLulus = nilai_absensi >= nilai_lulus
		cek_lulus        = cek_nilaiLulus && cek_absensiLulus
	)

	fmt.Println(cek_lulus)
}
