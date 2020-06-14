package main

import (
	"GoToolsForCosCustomizer/util"
	"fmt"
)

func main() {
	// out, err := tools.MovePartition("/dev/sdb", "1", "+10G")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(out)
	// tools.ExtendPartition("/dev/sdb", "1", 8408917)
	// start, err := tools.ReadPartitionStart("/dev/sdb", "1")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(start)
	err := util.ExtendOemPartition("/dev/sdb", "3", "1", "1G")
	if err != nil {
		fmt.Println(err.Error())
	}
}
