package main

import (
	"GoToolsForCosCustomizer/tools"
	"fmt"
)

func main() {
	out, err := tools.MovePartition("/dev/sdb", "1", "+1G")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(out)
}
