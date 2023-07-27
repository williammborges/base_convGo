package bases

import (
	"base_convGo/utils"
	"fmt"
	"math"
	"strconv"
)

const baseNumOctal = 8

func ConvertOctal(newBase Base, number string) {
	fmt.Println("Convertion to base " + newBase.String() + " the octal number " + number)

	var convertedNumber string
	switch newBase {
	case Binary:
		convertedNumber = octal_binary(number)
	case Octal:
		convertedNumber = number
	case Decimal:
		convertedNumber = octal_to_decimal(number)
	case Hexadecimal:
		convertedNumber = octal_to_hexadecimal(number)
	}

	fmt.Println("Converted number: " + convertedNumber)
}

func octal_binary(number string) string {
	var result string
	chunkSize := 1

	chunks := utils.DivideString(number, chunkSize)

	var numConverted string
	for _, v := range chunks {
		x, _ := strconv.Atoi(v)

		thirdDigit := x % 2
		x /= 2

		secondDigit := x % 2
		x /= 2

		firstDigit := x

		numConverted = numConverted + (strconv.Itoa(firstDigit) + strconv.Itoa(secondDigit) + strconv.Itoa(thirdDigit))
	}
	result += numConverted

	return result
}

func octal_to_decimal(number string) string {
	var result string
	chunkSize := 1

	chunks := utils.DivideString(number, chunkSize)

	var numConverted int
	for i, v := range chunks {
		x, _ := strconv.Atoi(v)
		numConverted = numConverted + (x * int(math.Pow(baseNumOctal, float64((len(number)-1)-i))))
	}
	result += strconv.Itoa(numConverted)

	return result
}

func octal_to_hexadecimal(number string) string {
	numToBin := octal_binary(number)
	result := binary_to_hexadecimal(numToBin)
	return result
}
