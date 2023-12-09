package hellos

import "fmt"

/*
	This is a package without a go.mod.
	To use it within this project, make sure
	this project is inside you go workspace folder
	which is declared $GOPATH
*/

func MyUtilityFunction() {
	fmt.Println("This is a utility function")
}
