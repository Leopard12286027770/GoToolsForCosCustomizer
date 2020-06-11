package main

import (
	"GoToolsForCosCustomizer/tools"
	"fmt"
)

func main() {
	// tools.Read_whole_disk()
	// tools.Read_disk_stat()
	fmt.Println("sdb1::")
	tools.Read_disk_size("/dev/sdb1")
	fmt.Println("sdb::")
	tools.Read_disk_size("/dev/sdb")

}
