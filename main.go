package main

import (
	functions "base_convGo/functions"
)

func main() {

	functions.Logo()

	actualBase := functions.SelectBase()
	newBase := functions.SelectBase()
	number := functions.SetNumber()

	functions.Convert(actualBase, newBase, number)

}
