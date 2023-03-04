package dccmd

import "os/exec"

const cmdName = "docker"

// Down brings down the docker compose environment and removes unused volumes
func Down() (string, error) {
	cmdArgs := []string{"compose", "down", "-v"}
	result, err := exec.Command(cmdName, cmdArgs...).CombinedOutput()

	return string(result), err
}

// Up brings up the docker compose environment in disconnected mode
func Up() (string, error) {
	cmdArgs := []string{"compose", "up", "-d"}
	result, err := exec.Command(cmdName, cmdArgs...).CombinedOutput()

	return string(result), err
}

// Build runs a no-cache build
func Build() (string, error) {
	cmdArgs := []string{"compose", "build", "--no-cache"}
	result, err := exec.Command(cmdName, cmdArgs...).CombinedOutput()

	return string(result), err
}

// PassThrough passes the supplied arguments directly to docker compose
func PassThrough(args []string) (string, error) {
	compose := []string{"compose"}
	cmdArgs := append(compose, args...)
	result, err := exec.Command(cmdName, cmdArgs...).CombinedOutput()

	return string(result), err
}
