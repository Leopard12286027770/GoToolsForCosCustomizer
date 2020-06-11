package tools

import (
	"os"
)

//Seek_disk_end use seek to find the size of disk
func Seek_disk_end(disk string) (int, error) {
	file, err := os.Open(disk)
	if Check(err) {
		return -1, err
	}
	// Whence is the point of reference for offset
	// 0 : Beginning of file
	// 1 : Current position
	// 2 : End of file
	pos, err := file.Seek(0, 2)
	if Check(err) {
		return -1, err
	}

	return int(pos), nil

}
