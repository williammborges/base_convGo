package utils

import "strconv"

func DivideString(str string, chunkSize int) []string {
	var chunks []string
	for i := 0; i < len(str); i += chunkSize {
		end := i + chunkSize
		if end > len(str) {
			end = len(str)
		}
		chunks = append(chunks, str[i:end])
	}
	return chunks
}

func VerifyMultiple(number string, multBy int) string {
	numberMultBy := number
	for len(numberMultBy)%multBy != 0 {
		numberMultBy = "0" + numberMultBy
	}

	return numberMultBy
}

func ChangeNumberToHexadecimal(number int) string {
	switch number {
	case 10:
		return "A"
	case 11:
		return "B"
	case 12:
		return "C"
	case 13:
		return "D"
	case 14:
		return "E"
	case 15:
		return "F"
	default:
		return strconv.Itoa(number)
	}
}

func ChangeNumberHexaToDecimal(number string) int {
	switch number {
	case "A":
		return 10
	case "B":
		return 11
	case "C":
		return 12
	case "D":
		return 13
	case "E":
		return 14
	case "F":
		return 15
	default:
		num, _ := strconv.Atoi(number)
		return num
	}
}
