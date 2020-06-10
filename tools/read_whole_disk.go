package tools

import (
        "fmt"
        "os"
)

func check(e error) bool {
        if e != nil {
                fmt.Println("ERROR!!!!")
                fmt.Println(e.Error())
                return true
        }
        return false
}

//Read_whole_disk to read whole disk output size in kB
func Read_whole_disk() {
        disk := "/dev/sdb1"
        file, err := os.Open(disk)
        if check(err) {
                return
        }
        buffer := make([]byte, 1024)
        num, err := file.Read(buffer)
        if check(err) {
                return
        }
        fmt.Println("READ completed")
        fmt.Println("num: ", num)
        sum := 1
        for {
                num, err = file.Read(buffer)
                sum++
                if check(err) {
						fmt.Println("total ", sum, " kB")
						file.Close()
                        return
                }
        }

}