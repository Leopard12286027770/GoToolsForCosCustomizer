package tools

import (
	"fmt"
	"os/exec"
)

//MovePartitionPlus move a partition to a start sector
// +XX(sector/G/M/K) or -XX(sector/G/M/K)
//for now it takes disk name like /dev/sda
//partition number like 2
//destination like 2048, +5G or -200M
func MovePartition(disk, partNum, dest string) error {
	cmd := "echo " + dest + " | sudo sfdisk --move-data " + disk + " -N " + partNum
	err := exec.Command("/bin/bash", "-c", cmd).Run()
	if Check(err) {
		fmt.Println("Error when moving partition")
		return err
	}
	return nil
}
