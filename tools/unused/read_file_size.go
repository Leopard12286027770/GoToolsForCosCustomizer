package tools

import (
	"fmt"
	"os"
)

//Read_file_size read a file size, cannot use read file to read a disk size
func Read_file_size(file string) {
	ff, err := os.Open(file)
	if Check(err) {
		return
	}
	stat, err := ff.Stat()
	if Check(err) {
		return
	}
	fmt.Println(stat.Size())
}
