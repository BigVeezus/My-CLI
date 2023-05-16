package main

import "fmt"

func commandAbout(cfg *config, args ...string) error{
	github := "https://github.com/bigveezus"
	// mailito := "elvis.osujic@gmail.com"
	fmt.Println("=========================================================")
	fmt.Println()
	fmt.Println("Heyy! I am Elvis.")
	fmt.Println("A software Engineer from Nigeria") 
	fmt.Println("I am available for open source backend projects.")
	fmt.Println("I write Javascript, Java and Go!")
	fmt.Println()

	fmt.Println("Check out my github,", github)
	fmt.Println()
	fmt.Println("=========================================================")
	return nil
}