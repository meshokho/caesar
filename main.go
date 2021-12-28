package main

import (
	"flag"
	"fmt"
)

// A simple test
func main() {
	text, shift, isEncode := getFlags()
	fmt.Println("Your text is:", text)

	switch shift {
	case 0:
		fmt.Println("I don't know the shift so i just make every possible, ok?")
		for i := 1; i <= 25; i++ {
			newString := cipher(text, i, isEncode)
			fmt.Println(i, "	- decoded:", newString)
		}
	default:
		newString := cipher(text, shift, isEncode)
		fmt.Println("encoded:", newString)
	}
}

func getFlags() (text string, shift int, encode bool) {
	flag.StringVar(&text, "text", "", "Text to decrypt")
	flag.IntVar(&shift, "shift", 0, "Shift")
	flag.BoolVar(&encode, "encode", true, "Encode = true, Decode = false")
	flag.Parse()

	return
}

func cipher(text string, shiftI int, isEncode bool) string {
	shift, offset := rune(shiftI), rune(26)

	runes := []rune(text)

	for index, char := range runes {
		switch isEncode {
		case false:
			if char >= 'a'+shift && char <= 'z' ||
				char >= 'A'+shift && char <= 'Z' {
				char = char - shift
			} else if char >= 'a' && char < 'a'+shift ||
				char >= 'A' && char < 'A'+shift {
				char = char - shift + offset
			}
		case true:
			if char >= 'a' && char <= 'z'-shift ||
				char >= 'A' && char <= 'Z'-shift {
				char = char + shift
			} else if char > 'z'-shift && char <= 'z' ||
				char > 'Z'-shift && char <= 'Z' {
				char = char + shift - offset
			}
		}

		runes[index] = char
	}

	return string(runes)
}
