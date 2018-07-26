package main

import (
	"os/exec"
	"syscall"
)

func main(){
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS
	}
}