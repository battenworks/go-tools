package main

import (
	"os"
	"strings"

	"github.com/battenworks/go-tools/common/v2/console"
	"github.com/battenworks/go-tools/tf/v2/tfcmd"
)

var version string = "built from source"

func main() {
	readable_version := "Version: " + strings.Replace(version, "tf_v", "", -1)

	if len(os.Args) > 1 {
		cmd := os.Args[1]

		switch cmd {
		case "version", "-v", "-version", "--version":
			console.Outln(readable_version)
		case "clean":
			workingDir := getWorkingDirectory()

			console.Outln("removing terraform cache")
			err := tfcmd.CleanTerraformCache(workingDir)
			if err != nil {
				console.Outln(err.Error())
				break
			}
			console.Greenln("terraform cache removed")

			console.Outln("initializing terraform")
			initResult, err := tfcmd.InitializeTerraform()
			console.Out(initResult)
			if err != nil {
				break
			}
		case "off":
			workingDir := getWorkingDirectory()

			err := tfcmd.Off(workingDir)
			if err != nil {
				console.Outln(err.Error())
				break
			}
		case "on":
			workingDir := getWorkingDirectory()

			err := tfcmd.On(workingDir)
			if err != nil {
				console.Outln(err.Error())
				break
			}
		case "test":
			console.Outln("validating config")
			result, err := tfcmd.PassThrough([]string{"validate"})
			if err != nil {
				console.Outln(err.Error())
				break
			}
			console.Out(result)
			console.Outln("")
		case "help", "-help", "--help":
			usage(readable_version)
		default:
			result, _ := tfcmd.PassThrough(os.Args[1:])
			console.Out(result)
		}
	} else {
		usage(readable_version)
	}
}

func getWorkingDirectory() string {
	scope, err := os.Getwd()
	if err != nil {
		console.Outln(err.Error())
		os.Exit(1)
	}

	workingDir, err := tfcmd.ValidateWorkingDirectory(scope)
	if err != nil {
		console.Outln(err.Error())
		os.Exit(1)
	}

	return workingDir
}

func usage(readable_version string) {
	console.Whiteln("Wrapper for the Terraform CLI")
	console.Whiteln("Provides some opinionated commands to help with Terraform CLI use")
	console.Whiteln("All other commands are passed directly to the Terraform CLI")
	console.Outln("")
	console.Whiteln(readable_version)
	console.Outln("")
	console.Whiteln("Usage: tf COMMAND")
	console.Outln("")
	console.Whiteln("commands:")
	console.Yellow("  clean")
	console.Whiteln("\t- Removes, then re-initializes, the Terraform cache of the current scope")
	console.Yellow("  off")
	console.Whiteln("\t- Adds the '.off' extension to all config files in the working directory")
	console.Whiteln("\t  Useful for preparing to destroy all resources in the current scope")
	console.Yellow("  on")
	console.Whiteln("\t- Removes the '.off' extension from all config files in the working directory")
	console.Whiteln("\t  Useful for preparing to re-create all resources in the current scope")
}
