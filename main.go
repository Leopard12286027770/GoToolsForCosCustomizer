package main

import "GoToolsForCosCustomizer/tools"

func main() {
	// tools.Read_whole_disk()
	tools.Read_disk_size("/dev/sdb1")
	tools.Read_disk_size("/dev/sdb")
}
