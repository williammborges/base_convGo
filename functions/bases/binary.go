package bases

import (
	"base_convGo/utils"
	"fmt"
	"strconv"
)

func ConvertBinary(newBase Base, number string) string {
	fmt.Println("Convertion to base " + newBase.String() + " the number " + number)

	var convertedNumber string
	switch newBase {
	case Binary:
		convertedNumber = number
	case Octal:
		convertedNumber = binary_to_octal(number)
	case Decimal:
		convertedNumber = strconv.Itoa(binary_to_decimal(number))
	case Hexadecimal:
		convertedNumber = binary_to_hexadecimal(number)
	}

	fmt.Println("Converted number: " + convertedNumber)

	return ""
}

func binary_to_octal(number string) string {
	var result string
	chunkSize := 3
	number = utils.VerifyMultiple(number, chunkSize)
	chunks := utils.DivideString(number, chunkSize)

	for _, n := range chunks {

		var chunksNum []string
		for i := 0; i < len(n); i += 1 {
			end := i + 1
			if end > len(n) {
				end = len(n)
			}
			chunksNum = append(chunks, n[i:end])
		}

		// n1 := strconv.Atoi(chunksNum[0]) * int(math.Pow(2, 2))
		// n2 := strconv.Atoi(chunksNum[1]) * int(math.Pow(2, 1))
		// n3 := strconv.Atoi(chunksNum[2]) * int(math.Pow(2, 0))

		fmt.Print(chunksNum[0], "\n")
		fmt.Print(chunksNum[1], "\n")
		fmt.Print(chunksNum[2], "\n")

		result = "" //result + strconv.Itoa((n1 + n2 + n3))
	}

	fmt.Print(chunks, "\n")

	return result
}

func divideString(number string, chunkSize int) {
	panic("unimplemented")
}

func binary_to_decimal(number string) int {
	return 0
}

func binary_to_hexadecimal(number string) string {
	return ""
}
