package gitkit

import (
	"os"
	"os/exec"
	"syscall"
)

func cleanUpProcessGroup(cmd *exec.Cmd) {
	if cmd == nil {
		return
	}

	process := cmd.Process
	if process != nil && process.Pid > 0 {
		p, err := os.FindProcess(-process.Pid)
		if err != nil {

		}
		err = p.Signal(syscall.SIGTERM)
		if err != nil {

		}
	}

	go cmd.Wait()
}
