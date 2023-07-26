package utils

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
