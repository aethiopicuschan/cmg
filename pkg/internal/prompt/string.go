package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"golang.org/x/term"
)

func PromptString(label, defaultValue string) string {
	cyan := color.New(color.FgCyan).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	scanner := bufio.NewScanner(os.Stdin)
	if defaultValue == "" {
		fmt.Printf("%s: ", cyan(label))
	} else {
		fmt.Printf("%s [%s]: ", cyan(label), yellow(defaultValue))
	}

	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())

	if input == "" {
		return defaultValue
	}
	return input
}

func PromptSecretString(label string) string {
	cyan := color.New(color.FgCyan).SprintFunc()
	fmt.Printf("%s: ", cyan(label))
	b, _ := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	return strings.TrimSpace(string(b))
}
