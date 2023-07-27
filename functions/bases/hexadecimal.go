package bases

import (
	"fmt"
	"math"
	"strconv"

	utils "github.com/williammborges/base_convGo/utils"
)

const baseNumHexa = 16

func ConvertHexadecimal(newBase Base, number string) {
	fmt.Println("Convertion to base " + newBase.String() + " the hexadecimal number " + number)

	var convertedNumber string
	switch newBase {
	case Binary:
		convertedNumber = hexadecimal_to_binary(number)
	case Octal:
		convertedNumber = hexadecimal_to_octal(number)
	case Decimal:
		convertedNumber = hexadecimal_to_decimal(number)
	case Hexadecimal:
		convertedNumber = number
	}

	fmt.Println("Converted number: " + convertedNumber)
}

func hexadecimal_to_binary(number string) string {
	var result string
	chunkSize := 1

	chunks := utils.DivideString(number, chunkSize)

	var numConverted string
	for _, v := range chunks {
		num := utils.ChangeNumberHexaToDecimal(v)

		fourthDigit := num % 2
		num /= 2

		thirdDigit := num % 2
		num /= 2

		secondDigit := num % 2
		num /= 2

		firstDigit := num

		numConverted = numConverted + (strconv.Itoa(firstDigit) + strconv.Itoa(secondDigit) + strconv.Itoa(thirdDigit) + strconv.Itoa(fourthDigit))
	}
	result += numConverted

	return result
}

func hexadecimal_to_octal(number string) string {
	num := hexadecimal_to_binary(number)
	result := binary_to_octal(num)
	return result
}

func hexadecimal_to_decimal(number string) string {
	var result string
	chunkSize := 1

	chunks := utils.DivideString(number, chunkSize)

	var numConverted int
	for i, v := range chunks {
		x := utils.ChangeNumberHexaToDecimal(v)
		numConverted = numConverted + (x * int(math.Pow(baseNumHexa, float64((len(number)-1)-i))))
	}
	result += strconv.Itoa(numConverted)

	return result
}
