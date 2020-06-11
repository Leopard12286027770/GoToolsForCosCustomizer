package main

func main() {
	// tools.Read_whole_disk("/dev/sdb1")
	// tools.Read_disk_size("/dev/sdb1")
	// tools.Read_disk_size("/dev/sdb")
	// tools.Read_disk_size("/dev/sda1")
	// tools.Read_disk_size("/dev/sda15")
	// tools.Read_disk_size("/dev/sda")
	tools.seek_disk_end("/dev/sdb1")
	tools.seek_disk_end("/dev/sdb")
}
