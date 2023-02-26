package tfcmd

import (
	"errors"
	"os"
	"path/filepath"
)

const cmdName = "terraform"

// ErrInvalidWorkingDirectory error constant.
var ErrInvalidWorkingDirectory = errors.New("invalid working directory: no backend.tf found")

// OffFileExtension is the file extension used to turn .tf files off and on.
var OffFileExtension = ".off"

// Executor is an interface to the struct in the command module
type Executor interface {
	Execute(cmdName string, cmdArgs ...string) ([]byte, error)
}

// ValidateWorkingDirectory determines if the supplied directory can be manipulated by this app.
func ValidateWorkingDirectory(dir string) (string, error) {
	if _, err := os.Stat(dir + "/backend.tf"); errors.Is(err, os.ErrNotExist) {
		return dir, ErrInvalidWorkingDirectory
	}

	return dir, nil
}

// CleanTerraformCache removes module cache and lock files from the supplied directory.
func CleanTerraformCache(dir string) error {
	err := os.RemoveAll(dir + "/.terraform")
	if err != nil {
		return err
	}

	if _, err := os.Stat(dir + "/.terraform.lock.hcl"); errors.Is(err, os.ErrNotExist) {
		return nil
	}

	err = os.Remove(dir + "/.terraform.lock.hcl")
	if err != nil {
		return err
	}

	return nil
}

// InitializeTerraform runs the Terraform init command.
func InitializeTerraform(executor Executor) (string, error) {
	cmdArgs := []string{"init"}
	result, err := executor.Execute(cmdName, cmdArgs...)

	return string(result), err
}

// Off adds a file extension to select Terraform config files,
// effectively turning them off for subsequent Terraform operations.
func Off(dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() && CanTurnFileOff(file.Name()) {
			filePath := dir + "/" + file.Name()

			err := os.Rename(filePath, filePath+OffFileExtension)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// CanTurnFileOff determines whether or not a file can have the "off" file extension.
func CanTurnFileOff(file string) bool {
	return file != "backend.tf" && file != "providers.tf" && filepath.Ext(file) == ".tf"
}

// On removes the file extensions that makes Terraform ignore config files,
// effectively turning them on for subsequent Terraform operations.
func On(dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() && CanTurnFileOn(file.Name()) {
			filePath := dir + "/" + file.Name()

			err := os.Rename(filePath, filePath[0:len(filePath)-len(OffFileExtension)])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// CanTurnFileOn determines whether or not a file can have the "off" extenstion removed.
func CanTurnFileOn(file string) bool {
	return filepath.Ext(file) == OffFileExtension
}

// PassThrough simply passes the commands to the Terraform binary, unmodified.
func PassThrough(executor Executor, cmdArgs []string) (string, error) {
	result, err := executor.Execute(cmdName, cmdArgs...)

	return string(result), err
}
