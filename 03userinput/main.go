package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Printf(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating to the web app")

	//commss ok // err ok
	input, err := reader.ReadString('\n')
	fmt.Println("Thanks for rating,", input)
	fmt.Printf("Type %T", input)
}