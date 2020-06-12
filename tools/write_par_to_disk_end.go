package tools

import (
	"fmt"
	"os"
)

const BUFFER_SIZE int64 = 4194304 //1MB

//Write_par_to_disk_end write the data in a partition to the end of the disk
func WriteParToDiskEnd(part, disk string) {
	partFile, err := os.Open(part)
	if Check(err) {
		return
	}
	diskFile, err := os.OpenFile(disk, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if Check(err) {
		return
	}
	defer partFile.Close()
	defer diskFile.Close()

	//seek to the end of the partition and disk
	// Whence is the point of reference for offset
	// 0 : Beginning of file
	// 1 : Current position
	// 2 : End of file
	partEnd, err := partFile.Seek(0, 2)
	if Check(err) {
		return
	}
	diskEnd, err := diskFile.Seek(-1073741824, 2) //last but 1 G
	if Check(err) {
		return
	}

	//check whether disk is large enough to hold the partition
	if diskEnd < partEnd {
		fmt.Println("no enough space in disk")
		return
	}
	diskStart := diskEnd - partEnd

	buffer := make([]byte, BUFFER_SIZE)
	for partEnd > 0 {
		offset := BUFFER_SIZE      //length of read and write
		if partEnd < BUFFER_SIZE { //no enough data for the buffer size
			offset = partEnd
		}
		//seek backwards
		partFile.Seek(-offset, 1)
		diskFile.Seek(-offset, 1)
		_, err = partFile.Read(buffer)
		if Check(err) {
			return
		}
		_, err = diskFile.Write(buffer[:offset])
		if Check(err) {
			return
		}
		partEnd -= offset
	}
	diskStart, err = diskFile.Seek(0, 1)
	if Check(err) {
		return
	}
	fmt.Println("Write completed. From ", part, " to the end of ", disk)
	fmt.Println("Start position on disk:", diskStart, "Byte")
}
