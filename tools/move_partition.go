package tools

import (
	"fmt"
	"os/exec"
)

//MovePartitionPlus move a partition to a start sector
// +XX(sector/G/M/K) or -XX(sector/G/M/K)
func MovePartition(disk, partNum, dest string) (string, error) {
	//echo '-1G,' | sudo sfdisk --move-data /dev/sdb -N 2
	cmd := "echo " + dest + " | sudo sfdisk --move-data " + disk + " -N " + partNum
	out, err := exec.Command("/bin/bash", "-c", cmd).Output()
	if Check(err) {
		fmt.Println("Error when moving partition")
		return "", err
	}
	return string(out), nil
}
