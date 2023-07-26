package bases

import (
	"fmt"
)

func ConvertDecimal(newBase Base, number string) string {
	fmt.Println("to base " + newBase.String() + " the number " + number)
	return ""
}
