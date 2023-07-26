package functions

import (
	baseConv "base_convGo/functions/bases"
	"fmt"
	"strconv"
)

func Logo() {
	fmt.Print(`

	▄▄▄▄    ▄▄▄        ██████ ▓█████     ▄████▄   ▒█████   ███▄    █ ██▒   █▓  ▄████  ▒█████  
	▓█████▄ ▒████▄    ▒██    ▒ ▓█   ▀    ▒██▀ ▀█  ▒██▒  ██▒ ██ ▀█   █▓██░   █▒ ██▒ ▀█▒▒██▒  ██▒
	▒██▒ ▄██▒██  ▀█▄  ░ ▓██▄   ▒███      ▒▓█    ▄ ▒██░  ██▒▓██  ▀█ ██▒▓██  █▒░▒██░▄▄▄░▒██░  ██▒
	▒██░█▀  ░██▄▄▄▄██   ▒   ██▒▒▓█  ▄    ▒▓▓▄ ▄██▒▒██   ██░▓██▒  ▐▌██▒ ▒██ █░░░▓█  ██▓▒██   ██░
	░▓█  ▀█▓ ▓█   ▓██▒▒██████▒▒░▒████▒   ▒ ▓███▀ ░░ ████▓▒░▒██░   ▓██░  ▒▀█░  ░▒▓███▀▒░ ████▓▒░
	░▒▓███▀▒ ▒▒   ▓▒█░▒ ▒▓▒ ▒ ░░░ ▒░ ░   ░ ░▒ ▒  ░░ ▒░▒░▒░ ░ ▒░   ▒ ▒   ░ ▐░   ░▒   ▒ ░ ▒░▒░▒░ 
	▒░▒   ░   ▒   ▒▒ ░░ ░▒  ░ ░ ░ ░  ░     ░  ▒     ░ ▒ ▒░ ░ ░░   ░ ▒░  ░ ░░    ░   ░   ░ ▒ ▒░ 
	░    ░   ░   ▒   ░  ░  ░     ░      ░        ░ ░ ░ ▒     ░   ░ ░     ░░  ░ ░   ░ ░ ░ ░ ▒  
	░            ░  ░      ░     ░  ░   ░ ░          ░ ░           ░      ░        ░     ░ ░  
		░                              ░                                ░                    
											
		
	By: SharedKernel																																																	  
	`)
}

func SelectBase() int {
	var base int

	for base == 0 {

		fmt.Println("Select the base to convert to:")
		fmt.Println(strconv.Itoa(baseConv.Binary.EnumIndex()) + ". " + baseConv.Binary.String())
		fmt.Println(strconv.Itoa(baseConv.Octal.EnumIndex()) + ". " + baseConv.Octal.String())
		fmt.Println(strconv.Itoa(baseConv.Decimal.EnumIndex()) + ". " + baseConv.Decimal.String())
		fmt.Println(strconv.Itoa(baseConv.Hexadecimal.EnumIndex()) + ". " + baseConv.Hexadecimal.String())

		fmt.Scanf("%d", &base)

		if base <= 0 || base > 4 {
			fmt.Println("Base to convert not found!")
			base = 0
		}
	}

	return base
}

func SetNumber() string {
	var number string
	fmt.Println("Enter the number to convert:")
	fmt.Scanf("%s", &number)

	return number
}

func Convert(actual_base int, new_base int, number string) {
	var convertedNumber string
	switch actual_base {
	case 1:
		convertedNumber = baseConv.ConvertBinary(baseConv.Base(new_base), number)
	case 2:
		convertedNumber = baseConv.ConvertOctal(baseConv.Base(new_base), number)
	case 3:
		convertedNumber = baseConv.ConvertDecimal(baseConv.Base(new_base), number)
	case 4:
		convertedNumber = baseConv.ConvertHexadecimal(baseConv.Base(new_base), number)
	default:
		convertedNumber = "Never converted!"
	}

	fmt.Println("Número convertido: " + convertedNumber)
	// i, _ := strconv.ParseInt(number, 0, 64)
	// var convertedNumber string
	// if base == 1 {
	// 	convertedNumber = strconv.FormatInt(i, 2)
	// } else if base == 2 {
	// 	convertedNumber = strconv.FormatInt(i, 8)
	// } else if base == 3 {
	// 	convertedNumber = strconv.FormatInt(i, 16)
	// }

	// // Print the converted number
	// fmt.Println("The converted number is:", convertedNumber)
}
