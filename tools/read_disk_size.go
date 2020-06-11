package tools

import (
	"fmt"
	"os"
)

//Read_disk_size read a disk size
func Read_disk_size(disk string) {
	file, err := os.Open(disk)
	if Check(err) {
		return
	}
	stat, err := file.Stat()
	if Check(err) {
		return
	}
	fmt.Println(stat.Size())
}
