package container

import (
	"os"
	"os/exec"
	"syscall"
)

// NewParentProcess : Create a paranet process
func NewParentProcess(tty bool, command string) *exec.Cmd {
	// initCommand RunContainerInitProcess
	args := []string{"init", command}
	cmd := exec.Command("/proc/self/exe", args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS |
			syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
	}
	if tty {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	return cmd
}

// Pipe : system pipe
func Pipe() (*os.File, *os.File, error) {
	read, write, err := os.Pipe()
	if err == nil {
		return read, write, nil
	}
	return nil, nil, err
}
