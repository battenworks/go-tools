package main

import (
	"os"

	"github.com/battenworks/go-tools/common/v2/console"
	"github.com/battenworks/go-tools/dc/v2/dccmd"
)

var version string = "built from source"

func main() {
	if len(os.Args) > 1 {
		cmd := os.Args[1]

		switch cmd {
		case "version", "-v", "-version", "--version":
			console.Outln("version: " + version)
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
			usage()
		default:
			result, _ := dccmd.PassThrough(os.Args[1:])
			console.Out(result)
		}
	}
}

func usage() {
	console.Whiteln("Wrapper for the Docker Compose CLI")
	console.Whiteln("Provides some opinionated commands to help with Docker Compose CLI use")
	console.Whiteln("All other commands are passed directly to the Docker Compose CLI")
	console.Outln("")
	console.Whiteln("Version: " + version)
	console.Outln("")
	console.Whiteln("Usage: dc COMMAND")
	console.Outln("")
	console.Whiteln("commands:")
	console.Yellow("  up")
	console.Whiteln("\t- Brings the docker-compose environment up in disconnected mode")
	console.Yellow("  down")
	console.Whiteln("\t- Brings down the docker-compose environment and removes all volumes")
	console.Yellow("  rebuild")
	console.Whiteln("\t- Brings down the docker-compose environment")
	console.Whiteln("\t  Removes all volumes")
	console.Whiteln("\t  Builds the docker-compose environment")
	console.Whiteln("\t  Brings the docker-compose environment up in disconnected mode")
}
