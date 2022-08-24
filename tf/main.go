package tf

import (
	"os"

	"github.com/battenworks/go-tools/common/v2/command"
	"github.com/battenworks/go-tools/common/v2/console"
)

// CommandExecutor is an interface to the struct in the command module
type CommandExecutor interface {
	Execute(cmdName string, cmdArgs ...string) ([]byte, error)
}

var executor CommandExecutor
var version string = "built from source"

func main() {
	executor = command.Executor{}

	if len(os.Args) > 1 {
		cmd := os.Args[1]

		switch cmd {
		case "version", "-v", "-version", "--version":
			console.Outln("version: " + version)
		case "clean":
			workingDir := getWorkingDirectory()

			console.Outln("removing terraform cache")
			err := CleanTerraformCache(workingDir)
			if err != nil {
				console.Outln(err.Error())
				break
			}
			console.Greenln("terraform cache removed")

			console.Outln("initializing terraform")
			initResult, err := InitializeTerraform(executor)
			console.Out(initResult)
			if err != nil {
				break
			}
		case "qplan":
			result := QuietPlan(executor)

			console.Out(result)
		case "off":
			workingDir := getWorkingDirectory()

			err := Off(workingDir)
			if err != nil {
				console.Outln(err.Error())
				break
			}
		case "on":
			workingDir := getWorkingDirectory()

			err := On(workingDir)
			if err != nil {
				console.Outln(err.Error())
				break
			}
		case "help", "-help", "--help":
			usage()
		default:
			result, _ := passThrough(executor, os.Args[1:])
			console.Out(result)
		}
	} else {
		usage()
	}
}

func getWorkingDirectory() string {
	scope, err := os.Getwd()
	if err != nil {
		console.Outln(err.Error())
		os.Exit(1)
	}

	workingDir, err := ValidateWorkingDirectory(scope)
	if err != nil {
		console.Outln(err.Error())
		os.Exit(1)
	}

	return workingDir
}

func usage() {
	console.Whiteln("Wrapper for the Terraform CLI")
	console.Whiteln("Provides some opinionated commands to help with Terraform CLI use")
	console.Whiteln("All other commands are passed directly to the Terraform CLI")
	console.Outln("")
	console.Whiteln("Version: " + version)
	console.Outln("")
	console.Whiteln("Usage: tf COMMAND")
	console.Outln("")
	console.Whiteln("commands:")
	console.Yellow("  clean")
	console.Whiteln("\t- Removes, then re-initializes, the Terraform cache of the current scope")
	console.Yellow("  qplan")
	console.Whiteln("\t- Calls terraform plan and hides drift output that results from the refresh stage of the plan")
	console.Yellow("  off")
	console.Whiteln("\t- Adds the '.off' extension to all config files in the working directory")
	console.Whiteln("\t  Useful for preparing to destroy all resources in the current scope")
	console.Yellow("  on")
	console.Whiteln("\t- Removes the '.off' extension from all config files in the working directory")
	console.Whiteln("\t  Useful for preparing to re-create all resources in the current scope")
}
