package tools

import (
	"errors"
	"os/exec"
	"strconv"
	"strings"
)

func ReadPartitionStart(disk, partNum string) (int, error) {
	//dump partition table to a file
	cmd := string("sfdisk --dump ") + disk
	table, err := exec.Command("/bin/bash", "-c", cmd).Output()
	if Check(err) {
		return -1, err
	}
	partName := disk + partNum
	lines := strings.Split(string(table), "\n")
	start := -1
	for _, line := range lines {
		if strings.Contains(line, partName) {
			ls := strings.Split(line, " ")
			mode := 0
			for _, word := range ls {
				switch mode {
				case 0: //looking for start sector
					if word == "start=" {
						mode = 1
					}
				case 1:
					if len(word) > 3 { //a valid sector number has at least 4 digits
						mode = 2
						start, err = strconv.Atoi(word[:len(word)-1]) //a comma at the end
						if Check(err) {
							return 0, err
						}
					}
				default:
					return -1, errors.New("Error: error in looking for partition")
				}
				if mode == 2 {
					break
				}
			}
			break
		}
	}
	if start == -1 {
		return -1, errors.New("Error: error in looking for partition")
	}
	return start, nil
}
