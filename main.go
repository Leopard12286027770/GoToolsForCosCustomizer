package main

import (
	"GoToolsForCosCustomizer/util"
	"fmt"
)

func main() {
	err := util.ExtendOemPartition("/dev/sdb", "3", "1", "1G")
	if err != nil {
		fmt.Println(err.Error())
	}
}
