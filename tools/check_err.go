package tools

import "fmt"

//Check check whether err is returned
func Check(e error) bool {
	if e != nil {
		fmt.Println("ERROR!!!!")
		fmt.Println(e.Error())
		return true
	}
	return false
}
