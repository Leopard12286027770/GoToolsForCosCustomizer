package tools

import (
	"fmt"
	"os"
)

//Read_whole_disk to read whole disk output size in B
func Read_whole_disk() {
	disk := "/dev/sdb1"
	file, err := os.Open(disk)
	if Check(err) {
		return
	}
	buffer := make([]byte, 1024)
	num, err := file.Read(buffer)
	if Check(err) {
		return
	}
	fmt.Println("READ completed")
	fmt.Println("num: ", num)
	sum := num
	for {
		num, err = file.Read(buffer)
		sum += num
		if num < 1024 {
			fmt.Println("total ", sum, " B")
			file.Close()
			return
		}
		if Check(err) {
			fmt.Println("total ", sum, " B")
			file.Close()
			return
		}
	}

}
