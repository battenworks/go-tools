package main

import (
	"os"
	"strings"

	"github.com/battenworks/go-tools/common/v2/console"
	"github.com/battenworks/go-tools/dc/v2/dccmd"
)

var version string = "built from source"

func main() {
	readable_version := "Version: " + strings.Replace(version, "dc_v", "", -1)

	if len(os.Args) > 1 {
		cmd := os.Args[1]

		switch cmd {
		case "version", "-v", "-version", "--version":
			console.Outln(readable_version)
		case "up":
			result, err := dccmd.Up()
			if err != nil {
				console.Outln(err.Error())
				break
			}
			console.Out(result)
			console.Outln("")
		case "down":
			result, err := dccmd.Down()
			if err != nil {
				console.Outln(err.Error())
				break
			}
			console.Out(result)
			console.Outln("")
		case "rebuild":
			result, err := dccmd.Down()
			if err != nil {
				console.Outln(err.Error())
				break
			}
			console.Out(result)
			console.Outln("")

			result, err = dccmd.Build()
			if err != nil {
				console.Outln(err.Error())
				break
			}
			console.Out(result)
			console.Outln("")

			result, err = dccmd.Up()
			if err != nil {
				console.Outln(err.Error())
				break
			}
			console.Out(result)
			console.Outln("")
		case "help", "-help", "--help":
			usage(readable_version)
		default:
			result, _ := dccmd.PassThrough(os.Args[1:])
			console.Out(result)
		}
	} else {
		usage(readable_version)
	}
}

func usage(readable_version string) {
	console.Whiteln("Wrapper for the Docker Compose CLI")
	console.Whiteln("Provides some opinionated commands to help with Docker Compose use")
	console.Whiteln("All commands not built into this tool are passed directly to Docker Compose")
	console.Outln("")
	console.Whiteln(readable_version)
	console.Outln("")
	console.Whiteln("Usage: dc COMMAND")
	console.Outln("")
	console.Whiteln("Commands:")
	console.Yellow("up")
	console.Whiteln("\t- Brings the docker-compose environment up in disconnected mode")
	console.Yellow("down")
	console.Whiteln("\t- Brings down the docker-compose environment and removes all volumes")
	console.Yellow("rebuild")
	console.Whiteln("\t- Brings down the docker-compose environment")
	console.Whiteln("\t  Removes all volumes")
	console.Whiteln("\t  Builds the docker-compose environment")
	console.Whiteln("\t  Brings the docker-compose environment up in disconnected mode")
}
