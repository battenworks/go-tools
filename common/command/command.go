package command

import (
	"os/exec"
)

type Executor struct{}

// Execute executes the given command with the given arguments.
func (ce Executor) Execute(cmdName string, cmdArgs ...string) ([]byte, error) {
	cmd := exec.Command(cmdName, cmdArgs...)
	return cmd.CombinedOutput()
}
