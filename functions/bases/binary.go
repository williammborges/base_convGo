package bases

import (
	"base_convGo/utils"
	"fmt"
	"math"
	"strconv"
)

const baseNumBinary = 2

func ConvertBinary(newBase Base, number string) {
	fmt.Println("Convertion to base " + newBase.String() + " the binary number " + number)

	var convertedNumber string
	switch newBase {
	case Binary:
		convertedNumber = number
	case Octal:
		convertedNumber = binary_to_octal(number)
	case Decimal:
		convertedNumber = binary_to_decimal(number)
	case Hexadecimal:
		convertedNumber = binary_to_hexadecimal(number)
	}

	fmt.Println("Converted number: " + convertedNumber)
}

func binary_to_octal(number string) string {
	var result string
	chunkSize := 3
	number = utils.VerifyMultiple(number, chunkSize)
	chunks := utils.DivideString(number, chunkSize)

	for _, n := range chunks {

		chunksNum := utils.DivideString(n, 1)
		var numConverted int
		for numConv := 0; numConv < chunkSize; numConv++ {
			x, _ := strconv.Atoi(chunksNum[numConv])
			numConverted = numConverted + (x * int(math.Pow(baseNumBinary, float64((chunkSize-1)-numConv))))
		}

		result += strconv.Itoa(numConverted)
	}

	return result
}

func binary_to_decimal(number string) string {
	var result string
	chunkSize := 1

	chunks := utils.DivideString(number, chunkSize)

	var numConverted int
	for i, v := range chunks {
		x, _ := strconv.Atoi(v)
		numConverted = numConverted + (x * int(math.Pow(baseNumBinary, float64((len(number)-1)-i))))
	}
	result += strconv.Itoa(numConverted)

	return result
}

func binary_to_hexadecimal(number string) string {
	var result string
	chunkSize := 4
	number = utils.VerifyMultiple(number, chunkSize)
	chunks := utils.DivideString(number, chunkSize)

	for _, n := range chunks {

		chunksNum := utils.DivideString(n, 1)
		var numConverted int
		for numConv := 0; numConv < chunkSize; numConv++ {
			x, _ := strconv.Atoi(chunksNum[numConv])
			numConverted = numConverted + (x * int(math.Pow(baseNumBinary, float64((chunkSize-1)-numConv))))
		}

		if numConverted > 9 {
			result += utils.ChangeNumberToHexadecimal(numConverted)
		} else {
			result += strconv.Itoa(numConverted)
		}

	}

	return result
}
