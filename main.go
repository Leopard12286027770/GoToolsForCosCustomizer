package main

import "GoToolsForCosCustomizer/tools"

func main() {
	// out, err := tools.MovePartition("/dev/sdb", "1", "-1G")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(out)
	tools.ExtendPartition("/dev/sdb", "1", 1000)

}
