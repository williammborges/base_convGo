package functions

import (
	"fmt"
	"strconv"

	baseConv "github.com/williammborges/base_convGo/functions/bases"
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
	switch actual_base {
	case 1:
		baseConv.ConvertBinary(baseConv.Base(new_base), number)
	case 2:
		baseConv.ConvertOctal(baseConv.Base(new_base), number)
	case 3:
		baseConv.ConvertDecimal(baseConv.Base(new_base), number)
	case 4:
		baseConv.ConvertHexadecimal(baseConv.Base(new_base), number)
	default:
		panic("oh no!")
	}
}
