package console_test

import (
	"testing"

	"github.com/battenworks/go-tools/common/v2/console"
)

func TestConsoleOut(t *testing.T) {
	console.Out("should be blue: ")
	console.Blue("blue")
	console.Outln("")

	console.Out("should be blue: ")
	console.Blueln("blue")

	console.Out("should be blue bold: ")
	console.BlueBold("blue bold")
	console.Outln("")

	console.Out("should be blue bold: ")
	console.BluelnBold("blue bold")

	console.Out("should be cyan: ")
	console.Cyan("cyan")
	console.Outln("")

	console.Out("should be cyan: ")
	console.Cyanln("cyan")

	console.Out("should be cyan bold: ")
	console.CyanBold("cyan bold")
	console.Outln("")

	console.Out("should be cyan bold: ")
	console.CyanlnBold("cyan bold")

	console.Out("should be green: ")
	console.Green("green")
	console.Outln("")

	console.Out("should be green: ")
	console.Greenln("green")

	console.Out("should be green bold: ")
	console.GreenBold("green bold")
	console.Outln("")

	console.Out("should be green bold: ")
	console.GreenlnBold("green bold")

	console.Out("should be magenta: ")
	console.Magenta("magenta")
	console.Outln("")

	console.Out("should be magenta: ")
	console.Magentaln("magenta")

	console.Out("should be magenta bold: ")
	console.MagentaBold("magenta bold")
	console.Outln("")

	console.Out("should be magenta bold: ")
	console.MagentalnBold("magenta bold")

	console.Out("should be red: ")
	console.Red("red")
	console.Outln("")

	console.Out("should be red: ")
	console.Redln("red")

	console.Out("should be red bold: ")
	console.RedBold("red bold")
	console.Outln("")

	console.Out("should be red bold: ")
	console.RedlnBold("red bold")

	console.Out("should be white: ")
	console.White("white")
	console.Outln("")

	console.Out("should be white: ")
	console.Whiteln("white")

	console.Out("should be white bold: ")
	console.WhiteBold("white bold")
	console.Outln("")

	console.Out("should be white bold: ")
	console.WhitelnBold("white bold")

	console.Out("should be yellow: ")
	console.Yellow("yellow")
	console.Outln("")

	console.Out("should be yellow: ")
	console.Yellowln("yellow")

	console.Out("should be yellow bold: ")
	console.YellowBold("yellow bold")
	console.Outln("")

	console.Out("should be yellow bold: ")
	console.YellowlnBold("yellow bold")
}
