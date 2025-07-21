package main

import "fmt"

func main() {
	const my_fav_fruit = "mango"

	const (
		database = "pendaftaran"
		port     = "5432"
	)

	fmt.Println(my_fav_fruit)
	fmt.Println(database)
	fmt.Println(port)
}
