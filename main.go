package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFromTerminal(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	str, err := reader.ReadString('\n')
	return str, err
}

func main() {
	word, _ := readFromTerminal("Enter a text")
	number, _ := readFromTerminal("Enter a number between -100 to 100")
	trimmedNumber := strings.Trim(number, "\n")
	intifiedNumber, err := strconv.Atoi(trimmedNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The Cipher Text is: ")
	encrypt(word, intifiedNumber)
}
