package main

import "GoToolsForCosCustomizer/tools"

func main() {
	// tools.Read_whole_disk()
	tools.Read_disk_size("/dev/sdb1")
	tools.Read_disk_size("/dev/sdb")
	tools.Read_disk_size("/dev/sda1")
	tools.Read_disk_size("/dev/sda15")
	tools.Read_disk_size("/dev/sda")
}
