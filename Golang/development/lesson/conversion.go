package main

import "fmt"

func main() {
	var name = "Danang"
	var find_word = name[1]
	// convert byte to string
	var convert_to_string = string(find_word)

	// before convert
	fmt.Println(find_word)
	// after convert
	fmt.Println(convert_to_string)
}
