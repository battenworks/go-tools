package dccmd

import "os/exec"

const cmdName = "docker"

func Down() (string, error) {
	cmdArgs := []string{"compose", "down", "-v"}
	result, err := exec.Command(cmdName, cmdArgs...).CombinedOutput()

	return string(result), err
}

func Up() (string, error) {
	cmdArgs := []string{"compose", "up", "-d"}
	result, err := exec.Command(cmdName, cmdArgs...).CombinedOutput()

	return string(result), err
}

func Build() (string, error) {
	cmdArgs := []string{"compose", "build"}
	result, err := exec.Command(cmdName, cmdArgs...).CombinedOutput()

	return string(result), err
}

func PassThrough(cmdArgs []string) (string, error) {
	result, err := exec.Command(cmdName, cmdArgs...).CombinedOutput()

	return string(result), err
}
