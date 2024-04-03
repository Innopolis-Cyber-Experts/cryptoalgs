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

func ToHex(data []byte) (result string) {
	for i := range len(data) {
		higherValue := data[i] / 0x10
		lowerValue := data[i] % 0x10

		var higherDigit rune
		var lowerDigit rune

		if higherValue <= 9 {
			higherDigit = rune(int('0') + int(higherValue))
		} else {
			higherDigit = rune(int('a') + int(higherValue) - 10)
		}

		if lowerValue <= 9 {
			lowerDigit = rune(int('0') + int(lowerValue))
		} else {
			lowerDigit = rune(int('a') + int(lowerValue) - 10)
		}

		result += string(higherDigit) + string(lowerDigit)
	}
	return
}

func Xor(data []byte, key []byte) (result []byte) {
	for i := range len(data) {
		result = append(result, data[i]^key[i])
	}
	return
}

var toXor string
var key string

func init() {
	flag.StringVar(&toXor, "s", "", "hex string to be xored with key")
	flag.StringVar(&key, "k", "", "key as hex string")
}

func main() {
	flag.Parse()

	toXorData, err := FromHex(toXor)
	if err != nil {
		fmt.Errorf("error: %s", err)
		return
	}

	keyData, err := FromHex(key)
	if err != nil {
		fmt.Errorf("error: %s", err)
		return
	}

	encryptedData := Xor(toXorData, keyData)

	encrypted := ToHex(encryptedData)

	fmt.Println(encrypted)
}
