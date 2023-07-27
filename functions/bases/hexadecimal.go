package bases

import (
	"fmt"
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
	panic("unimplemented")
}

func hexadecimal_to_octal(number string) string {
	panic("unimplemented")
}

func hexadecimal_to_decimal(number string) string {
	panic("unimplemented")
}
