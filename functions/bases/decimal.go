package bases

import (
	"fmt"
	"strconv"

	utils "github.com/williammborges/base_convGo/utils"
)

func ConvertDecimal(newBase Base, number string) {
	fmt.Println("Convertion to base " + newBase.String() + " the decimal number " + number)

	var convertedNumber string
	switch newBase {
	case Binary:
		convertedNumber = decimal_to_binary(number)
	case Octal:
		convertedNumber = decimal_to_octal(number)
	case Decimal:
		convertedNumber = number
	case Hexadecimal:
		convertedNumber = decimal_to_hexadecimal(number)
	}

	fmt.Println("Converted number: " + convertedNumber)
}

func decimal_to_binary(number string) string {
	var result string
	num, _ := strconv.Atoi(number)

	if num == 0 {
		return "0"
	}

	for num > 0 {
		rest := num % baseNumBinary
		quotient := num / baseNumBinary
		result = strconv.Itoa(rest) + result
		num = quotient
	}

	return result
}

func decimal_to_octal(number string) string {
	var result string
	num, _ := strconv.Atoi(number)

	if num == 0 {
		return "0"
	}

	for num > 0 {
		rest := num % baseNumOctal
		quotient := num / baseNumOctal
		result = strconv.Itoa(rest) + result
		num = quotient
	}

	return result
}

func decimal_to_hexadecimal(number string) string {
	var result string
	num, _ := strconv.Atoi(number)

	if num == 0 {
		return "0"
	}

	for num > 0 {
		rest := num % baseNumHexa
		quotient := num / baseNumHexa
		utils.ChangeNumberToHexadecimal(rest)
		result = utils.ChangeNumberToHexadecimal(rest) + result
		num = quotient
	}

	return result
}
