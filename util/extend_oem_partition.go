package util

import (
	"GoToolsForCosCustomizer/tools"
	"errors"
	"fmt"
	"strconv"
)

//ExtendOemPartition moves stateful partition towards the end of the disk
//Then move oem partition to the original place of the stateful partition
//Finally resize the oem partition to 1 sector before the new stateful partition
//oemSize can be the number of sectors (without unit) or size like "3G"
func ExtendOemPartition(disk, statePartNum, oemPartNum, oemSize string) error {
	//read original size of OEM partition
	oriOemSize, err := tools.ReadPartitionSize(disk, oemPartNum)
	if err != nil {
		return err
	}
	oriOemSizeBytes := oriOemSize * 512 //change unit to bytes
	newOemSizeBytes, err := tools.ConvertSizeToBytes(oemSize)

	if err != nil {
		return err
	}

	if newOemSizeBytes <= oriOemSizeBytes {
		return errors.New("Error: oemSize: " + strconv.Itoa(newOemSizeBytes) + " bytes is not larger than the original OEM partition size: " + strconv.Itoa(oriOemSizeBytes) + " bytes")
	}

	//record the original start sector of the stateful partition
	oriStartSector, err := tools.ReadPartitionStart(disk, statePartNum)
	if err != nil {
		return err
	}

	//move the stateful partition
	err = tools.MovePartition(disk, statePartNum, "+"+oemSize)
	if err != nil {
		return err
	}

	//move oem partition to the original start sector of the stateful partition
	err = tools.MovePartition(disk, oemPartNum, strconv.Itoa(oriStartSector))
	if err != nil {
		return err
	}

	//record the new start sector of the stateful partition
	newStartSector, err := tools.ReadPartitionStart(disk, statePartNum)
	if err != nil {

		return err
	}

	//extend the oem partition
	err = tools.ExtendPartition(disk, oemPartNum, newStartSector-1)
	if err != nil {
		return err
	}
	fmt.Printf("\nCompleted extending OEM partition\n")
	return nil
}
