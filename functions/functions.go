package functions

import (
	"fmt"
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
	bases := GetBases()

	for base == 0 {

		fmt.Println("Select the base to convert to:")
		for key, value := range bases {
			fmt.Println((key + 1), ". ", value)
		}

		fmt.Scanf("%d", &base)

		if base <= 0 || base > len(bases) {
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

func GetBases() [4]string {
	bases := [4]string{
		0: "Binary",
		1: "Octal",
		2: "Decimal",
		3: "Hexadecimal",
	}
	return bases
}
