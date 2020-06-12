package tools

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

//ExtendPartition extends a partition to a specific end sector
func ExtendPartition(disk, partNum string, end int) {
	// cmd := string("sfdisk --dump ") + disk + " > extend_partition_tmp"
	// exec.Command("/bin/bash", "-c", cmd).Run()
	partName := disk + partNum
	editPartitionTableFile("aaa", partName, 6666665)
}

//change partition table file to extend partition
func editPartitionTableFile(fileName, partName string, end int) error {
	in, err := ioutil.ReadFile(fileName)
	if Check(err) {
		return err
	}

	lines := strings.Split(string(in), "\n")
	have := false //whether has valid information about the partition
	for i, line := range lines {
		if strings.Contains(line, partName) {
			ls := strings.Split(line, " ")
			mode := 0
			start := -1
			for j, word := range ls {
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
							return err
						}
						fmt.Println("start:", start)
					}
				case 2:
					if word == "size=" {
						mode = 3
					}
				case 3:
					if len(word) > 3 { //a valid sector number has at least 4 digits

						size, err := strconv.Atoi(word[:len(word)-1]) //a comma at the end
						if Check(err) {
							return err
						}
						fmt.Println("size:", size)
						if end-start+1 <= size {
							return errors.New("Error: new size is maller than the original size!")
						}
						have = true
						ls[j] = strconv.Itoa(end+1-start) + ","
					}
				default:
					return errors.New("error in looking for partition")
				}
				if have {
					break
				}
			}
			if have {
				lines[i] = strings.Join(ls, " ")
			}
			break
		}
	}
	if !have {
		return errors.New("Error: Partition not found!")
	}
	changed := strings.Join(lines, "\n")
	err = ioutil.WriteFile(fileName, []byte(changed), 0644)
	if Check(err) {
		return err
	}
	fmt.Println(partName, "extension completed")
	return nil
}
