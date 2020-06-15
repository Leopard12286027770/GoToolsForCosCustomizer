package tools

import "fmt"

//Check check whether err is returned and print error message
func Check(e error, msg string) bool {
	if e != nil {
		fmt.Printf("ERROR!!!!\n %s\n %s\n\n", msg, e.Error())
		return true
	}
	return false
}
