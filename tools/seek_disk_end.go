package tools

import (
	"fmt"
	"os"
)

//Seek_disk_end use seek to find the size of disk
func Seek_disk_end(disk string) {
	file, err := os.Open(disk)
	if Check(err) {
		return
	}
	// Whence is the point of reference for offset
	// 0 : Beginning of file
	// 1 : Current position
	// 2 : End of file
	pos, err := file.Seek(0, 2)
	if Check(err) {
		return
	}
	fmt.Println(disk, ":")
	fmt.Println("Disk end position: ", pos)

}
