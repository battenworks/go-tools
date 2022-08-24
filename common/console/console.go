package console

import (
	"fmt"
)

const (
	format           = "%s%s%s"
	colorBlue        = "\x1b[34m"
	colorBlueBold    = "\x1b[34;1m"
	colorCyan        = "\x1b[36m"
	colorCyanBold    = "\x1b[36;1m"
	colorGreen       = "\x1b[32m"
	colorGreenBold   = "\x1b[32;1m"
	colorMagenta     = "\x1b[35m"
	colorMagentaBold = "\x1b[35;1m"
	colorRed         = "\x1b[31m"
	colorRedBold     = "\x1b[31;1m"
	colorWhite       = "\x1b[37m"
	colorWhiteBold   = "\x1b[37;1m"
	colorYellow      = "\x1b[33m"
	colorYellowBold  = "\x1b[33;1m"
	colorNone        = "\x1b[0m"
)

// Blue prints text in blue.
func Blue(text string) {
	fmt.Printf(format, colorBlue, text, colorNone)
}

// Blueln prints text in blue followed by a newline character.
func Blueln(text string) {
	fmt.Printf(format, colorBlue, text, colorNone+"\n")
}

// BlueBold prints bold text in blue.
func BlueBold(text string) {
	fmt.Printf(format, colorBlueBold, text, colorNone)
}

// BluelnBold prints bold text in blue followed by a newline character.
func BluelnBold(text string) {
	fmt.Printf(format, colorBlueBold, text, colorNone+"\n")
}

// Cyan prints text in cyan.
func Cyan(text string) {
	fmt.Printf(format, colorCyan, text, colorNone)
}

// Cyanln prints text in cyan followed by a newline character.
func Cyanln(text string) {
	fmt.Printf(format, colorCyan, text, colorNone+"\n")
}

// CyanBold prints bold text in cyan.
func CyanBold(text string) {
	fmt.Printf(format, colorCyanBold, text, colorNone)
}

// CyanlnBold prints bold text in cyan followed by a newline character.
func CyanlnBold(text string) {
	fmt.Printf(format, colorCyanBold, text, colorNone+"\n")
}

// Green prints text in green.
func Green(text string) {
	fmt.Printf(format, colorGreen, text, colorNone)
}

// Greenln prints text in green followed by a newline character.
func Greenln(text string) {
	fmt.Printf(format, colorGreen, text, colorNone+"\n")
}

// GreenBold prints bold text in green.
func GreenBold(text string) {
	fmt.Printf(format, colorGreenBold, text, colorNone)
}

// GreenlnBold prints bold text in green followed by a newline character.
func GreenlnBold(text string) {
	fmt.Printf(format, colorGreenBold, text, colorNone+"\n")
}

// Magenta prints text in magenta.
func Magenta(text string) {
	fmt.Printf(format, colorMagenta, text, colorNone)
}

// Magentaln prints text in magenta followed by a newline character.
func Magentaln(text string) {
	fmt.Printf(format, colorMagenta, text, colorNone+"\n")
}

// MagentaBold prints bold text in magenta.
func MagentaBold(text string) {
	fmt.Printf(format, colorMagentaBold, text, colorNone)
}

// MagentalnBold prints bold text in magenta followed by a newline character.
func MagentalnBold(text string) {
	fmt.Printf(format, colorMagentaBold, text, colorNone+"\n")
}

// Out prints text with no color modification.
func Out(text string) {
	fmt.Printf(format, colorNone, text, colorNone)
}

// Outln prints text with no color modification followed by a newline character.
func Outln(text string) {
	fmt.Printf(format, colorNone, text, colorNone+"\n")
}

// Red prints text in red.
func Red(text string) {
	fmt.Printf(format, colorRed, text, colorNone)
}

// Redln prints text in red followed by a newline character.
func Redln(text string) {
	fmt.Printf(format, colorRed, text, colorNone+"\n")
}

// RedBold prints bold text in red.
func RedBold(text string) {
	fmt.Printf(format, colorRedBold, text, colorNone)
}

// RedlnBold prints bold text in red followed by a newline character.
func RedlnBold(text string) {
	fmt.Printf(format, colorRedBold, text, colorNone+"\n")
}

// White prints text in white.
func White(text string) {
	fmt.Printf(format, colorWhite, text, colorNone)
}

// Whiteln prints text in white followed by a newline character.
func Whiteln(text string) {
	fmt.Printf(format, colorWhite, text, colorNone+"\n")
}

// WhiteBold prints bold text in white.
func WhiteBold(text string) {
	fmt.Printf(format, colorWhiteBold, text, colorNone)
}

// WhitelnBold prints bold text in white followed by a newline character.
func WhitelnBold(text string) {
	fmt.Printf(format, colorWhiteBold, text, colorNone+"\n")
}

// Yellow prints text in yellow.
func Yellow(text string) {
	fmt.Printf(format, colorYellow, text, colorNone)
}

// Yellowln prints text in yellow followed by a newline character.
func Yellowln(text string) {
	fmt.Printf(format, colorYellow, text, colorNone+"\n")
}

// YellowBold prints bold text in yellow.
func YellowBold(text string) {
	fmt.Printf(format, colorYellowBold, text, colorNone)
}

// YellowlnBold prints bold text in yellow followed by a newline character.
func YellowlnBold(text string) {
	fmt.Printf(format, colorYellowBold, text, colorNone+"\n")
}
