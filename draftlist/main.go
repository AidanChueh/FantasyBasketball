package main

import "fmt"

func main() {
	fmt.Println("Draft day boyz")
	s := fileToSlice("playerlist")
	fmt.Println(s)
	fmt.Println(s[0])
	fmt.Println(s[1])
}
