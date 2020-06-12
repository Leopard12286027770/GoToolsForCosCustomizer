package tools

import (
	"fmt"
	"os/exec"
)

//ExtendPartition extends a partition to a specific end sector
func ExtendPartition(disk, partNum string, end int) {
	cmd := string("sfdisk --dump ") + disk + " > extend_partition_tmp"
	exec.Command("/bin/bash", "-c", cmd).Run()
	part := disk + partNum
	fmt.Println(part)
}
