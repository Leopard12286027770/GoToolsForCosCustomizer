package tools

import (
	"os"
	"os/exec"
)

//Run a command and output to stdout
func ExecCmdToStdout(cmdLine string) error {
	cmd := exec.Command("/bin/bash", "-c", cmdLine)
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	return err
}
