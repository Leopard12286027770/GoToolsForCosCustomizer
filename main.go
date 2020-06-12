package main

import "GoToolsForCosCustomizer/tools"

func main() {
	tools.WriteParToDiskEnd("/dev/sdb1", "/dev/sdb")
	tools.CheckData("/dev/sdb", "/dev/sdb1")
}
