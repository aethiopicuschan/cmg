package colors

import "github.com/fatih/color"

func Cyan(s string) string {
	return color.New(color.FgCyan).SprintFunc()(s)
}

func Green(s string) string {
	return color.New(color.FgGreen, color.Bold).SprintFunc()(s)
}

func Red(s string) string {
	return color.New(color.FgRed).SprintFunc()(s)
}

func Yellow(s string) string {
	return color.New(color.FgYellow).SprintFunc()(s)
}

func Magenta(s string) string {
	return color.New(color.FgMagenta).SprintFunc()(s)
}

func Blue(s string) string {
	return color.New(color.FgBlue).SprintFunc()(s)
}
