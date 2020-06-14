package main

import "GoToolsForCosCustomizer/tools"

//ExtendOemPartition moves stateful partition towards the end of the disk
//Then move oem partition to the original place of the stateful partition
//Finally resize the oem partition to 1 sector before the new stateful partition
//oemSize can be the number of sectors (without unit) or size like "3G"
func ExtendOemPartition(disk, statePartNum, oemPartNum, oemSize string) error {
	//record the original start sector of the stateful partition
	oriStartSector, err := tools.ReadPartitionStart(disk, statePartNum)
	if err != nil {

		return err
	}

	//move the stateful partition
	err = tools.MovePartition(disk, statePartNum, "+"+oemSize)
	// if err
	return nil
}
