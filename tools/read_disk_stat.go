package tools

import (
	"fmt"
	"os"
	"syscall"
)

//Read_disk_stat read disk status
func Read_disk_stat() {
	var stat syscall.Statfs_t
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	syscall.Statfs(wd, &stat)

	// Available blocks * size per block = available space in bytes
	fmt.Println(stat)
}
