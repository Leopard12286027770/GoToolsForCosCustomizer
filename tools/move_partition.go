package tools

import (
	"fmt"
)

//MovePartitionPlus move a partition to a start sector
// +XX(sector/G/M/K) or -XX(sector/G/M/K)
//for now it takes disk name like /dev/sda
//partition number like 2
//destination like 2048, +5G or -200M
func MovePartition(disk, partNum, dest string) error {
	cmd := "echo " + dest + " | sudo sfdisk --move-data " + disk + " -N " + partNum
	err := ExecCmdToStdout(cmd)
	if Check(err, cmd) {
		return err
	}
	fmt.Printf("\nCompleted moving %s \n\n", (disk + partNum))
	return nil
}
