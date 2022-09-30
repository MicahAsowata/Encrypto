package main

import (
	"fmt"
)

func contains(stringSlice []string, text string) (bool, int) {
	for index, word := range stringSlice {

		if word == text {
			return true, int(index)
		}
	}
	return false, -1
}

func encrypt(plainText string, key int) {
	alphabet := []string{"A", "a", "B", "b", "C", "c", "D", "d", "E", "e", "F", "f", "G", "g", "H", "h", "I", "i", "J", "j", "K", "k", "L", "l", "M", "m", "N", "n", "O", "o", "P", "p", "Q", "q", "R", "r", "S", "s", "T", "t", "U", "u", "V", "v", "W", "w", "X", "x", "Y", "y", "Z", "z"}

	slicedText := []rune(plainText)

	for _, v := range slicedText {
		stringifiedSlicedText := string(v)
		checkPresence, _ := contains(alphabet, stringifiedSlicedText)
		if checkPresence {
			_, letter := contains(alphabet, stringifiedSlicedText)
			encryptSteps := letter + key
			if encryptSteps >= len(alphabet) {
				A := encryptSteps - len(alphabet)
				if A >= len(alphabet) {
					B := A - len(alphabet)
					fmt.Print(alphabet[B])
				} else {
					fmt.Print(alphabet[A])
				}
			} else if encryptSteps <= 0 {
				C := encryptSteps + len(alphabet)
				if C <= 0 {
					D := C + len(alphabet)
					fmt.Print(alphabet[D])
				} else {
					fmt.Print(alphabet[C])
				}
			} else {
				fmt.Print(alphabet[encryptSteps])
			}
		} else {
			if stringifiedSlicedText == " " {
				fmt.Print(stringifiedSlicedText)
			} else {
				fmt.Print(".")
			}
		}
	}
	fmt.Println()
}
