package main

import (
	"fmt"
	"os"
	"time"
	"unicode"
)

const Reset = "\033[0m"
const Red = "\033[31m"
const Green = "\033[32m"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You forgot to type parameter. Example:")
		fmt.Println("./check-ascii <your input>")
		return
	}

	fmt.Println("--------------------")
	fmt.Println("Started   :", time.Now().Format(time.RFC1123))

	input := os.Args[1]
	output := ""
	nonASCIICount := 0

	fmt.Println("Your input:", input)

	for _, c := range input {
		// Non-ASCII
		if c > unicode.MaxASCII {
			output += Red + string(c) + Reset
			nonASCIICount++
		} else {
			output += string(c)
		}
	}

	fmt.Println("Output    :", output)

	if nonASCIICount > 0 {
		fmt.Printf("Found %s%d%s Non-ASCII characters.\n", Red, nonASCIICount, Reset)
		fmt.Printf("%sNon-ASCII%s\n", Red, Reset)
	} else {
		fmt.Printf("Found %s%d%s Non-ASCII characters.\n", Green, 0, Reset)
		fmt.Printf("%sASCII%s\n", Green, Reset)
	}

	fmt.Println("Finished  :", time.Now().Format(time.RFC1123))
	fmt.Println("--------------------")
}
