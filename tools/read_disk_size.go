package tools

import (
	"fmt"
	"syscall"
)

//Read_disk_size read disk size
func Read_disk_size(disk string) {
	var stat syscall.Statfs_t
	err := syscall.Statfs(disk, &stat)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Available blocks * size per block = available space in bytes
	size := stat.Blocks * uint64(stat.Bsize)
	ava := stat.Bavail * uint64(stat.Bsize)
	fmt.Println("Disk: ", disk)

	fmt.Println("Disk size: ", size)
	fmt.Println("Disk available: ", ava)
	fmt.Println("Disk used: ", size-ava)
	fmt.Println()

}
