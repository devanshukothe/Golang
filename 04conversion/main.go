package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Enter input: ")
	reader := bufio.NewReader(os.Stdin)
   input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for reading", input)
    
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating:", numRating+1)
	}
}