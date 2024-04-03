package main

import (
	"errors"
	"flag"
	"fmt"
)

var hexValues = map[byte]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'a': 10,
	'b': 11,
	'c': 12,
	'd': 13,
	'e': 14,
	'f': 15,
}

func FromHex(hexString string) ([]byte, error) {
	if len(hexString)%2 != 0 {
		return nil, errors.New("hex string has not odd length")
	}

	var data []byte

	for pos := 0; pos < len(hexString); pos += 2 {
		higherDigit := hexString[pos]
		lowerDigit := hexString[pos+1]

		value := hexValues[lowerDigit] + hexValues[higherDigit]*0x10

		data = append(data, byte(value))
	}

	return data, nil
}

var alphabet []string = []string{
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
	"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
	"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "+", "/",
}

func ToBase64(data []byte) string {
	var encoded string

	for pos := 0; pos < len(data)/3*3; pos += 3 {
		byteValues := []byte{
			data[pos] & 0b1111100 >> 2,
			(data[pos] & 0b00000011 << 4) | (data[pos+1] & 0b11110000 >> 4),
			(data[pos+1] & 0b00001111 << 2) | (data[pos+2] & 0b11000000 >> 6),
			(data[pos+2] & 0b00111111),
		}

		encodedBytes := alphabet[byteValues[0]] + alphabet[byteValues[1]] + alphabet[byteValues[2]] + alphabet[byteValues[3]]

		encoded += encodedBytes
	}

	if len(data)%3 == 1 {
		pos := len(data) - 1
		byteValues := []byte{
			data[pos] & 0b1111100 >> 2,
			(data[pos] & 0b00000011 << 4),
		}

		encodedBytes := alphabet[byteValues[0]] + alphabet[byteValues[1]] + "=="

		encoded += encodedBytes
	} else if len(data)%3 == 2 {
		pos := len(data) - 2
		byteValues := []byte{
			data[pos] & 0b1111100 >> 2,
			(data[pos] & 0b00000011 << 4) | (data[pos+1] & 0b11110000 >> 4),
			(data[pos+1] & 0b00001111 << 2),
		}

		encodedBytes := alphabet[byteValues[0]] + alphabet[byteValues[1]] + alphabet[byteValues[2]] + "="

		encoded += encodedBytes
	}

	return encoded
}

var toTranslate string

func init() {
	flag.StringVar(&toTranslate, "s", "", "hex string to convert to base64")
}

func main() {
	flag.Parse()

	data, err := FromHex(toTranslate)
	if err != nil {
		fmt.Errorf("error: %s", err)
		return
	}

	encoded := ToBase64(data)
	fmt.Println(encoded)
}
