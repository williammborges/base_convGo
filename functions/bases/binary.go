package bases

import (
	"base_convGo/utils"
	"fmt"
	"math"
	"strconv"
)

const baseNum = 2

func ConvertBinary(newBase Base, number string) {
	fmt.Println("Convertion to base " + newBase.String() + " the number " + number)

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
			numConverted = numConverted + (x * int(math.Pow(baseNum, float64((chunkSize-1)-numConv))))
		}

		result += strconv.Itoa(numConverted)
	}

	return result
}

func binary_to_decimal(number string) string {
	var result string
	chunkSize := 1
	// number = utils.VerifyMultiple(number, chunkSize)
	chunks := utils.DivideString(number, chunkSize)

	var numConverted int
	for i, v := range chunks {
		x, _ := strconv.Atoi(v)
		numConverted = numConverted + (x * int(math.Pow(baseNum, float64((len(number)-1)-i))))
	}
	result += strconv.Itoa(numConverted)

	// for _, n := range chunks {

	// 	chunksNum := utils.DivideString(n, 1)
	// 	n1, _ := strconv.Atoi(chunksNum[0])
	// 	n2, _ := strconv.Atoi(chunksNum[1])
	// 	n3, _ := strconv.Atoi(chunksNum[2])

	// 	n1 = n1 * int(math.Pow(2, 2))
	// 	n2 = n2 * int(math.Pow(2, 1))
	// 	n3 = n3 * int(math.Pow(2, 0))

	// 	result = result + strconv.Itoa(n1+n2+n3)
	// }

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
			numConverted = numConverted + (x * int(math.Pow(baseNum, float64((chunkSize-1)-numConv))))
		}

		if numConverted > 9 {
			result += utils.ChangeNumberToDecimal(numConverted)
		} else {
			result += strconv.Itoa(numConverted)
		}

	}

	return result
}
