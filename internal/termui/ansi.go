package termui

import (
	"fmt"
	"strings"
)

const (
	CheckMark string = "\u2714"
	XMark     string = "\u2715"
)

type Color string

const (
	RedColor     Color = "31"
	GreenColor   Color = "32"
	YellowColor  Color = "33"
	BlueColor    Color = "34"
	MagentaColor Color = "35"
	CyanColor    Color = "36"
	GreyColor    Color = "90"

	BrightWhiteColor Color = "37;1"
)

func Red(text string) string {
	return Colorize(text, RedColor)
}

func Green(text string) string {
	return Colorize(text, GreenColor)
}

func Magenta(text string) string {
	return Colorize(text, MagentaColor)
}

func BrightWhite(text string) string {
	return Colorize(text, BrightWhiteColor)
}

func Grey(text string) string {
	return Colorize(text, GreyColor)
}

func Colorize(text string, color Color) string {
	sb := strings.Builder{}

	// set color
	fmt.Fprintf(&sb, "\u001B[%sm", color)

	fmt.Fprintf(&sb, text)

	// unset - reset
	fmt.Fprintf(&sb, "\u001B[0m")

	return sb.String()
}

// Print Hyperlink via OSC 8 ansi sequence.
// The syntax is: 'OSC 8 ; params ; url ST text OSC 8 ; ; ST'
// for more info see https://gist.github.com/egmontkob/eb114294efbcd5adb1944c9f3cb5feda
func Hyperlink(name, url string) string {
	return fmt.Sprintf("\u001B]8;%s;%s\u001B\\%s\u001B]8;;\u001B\\", "", url, name)
}
