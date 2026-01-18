package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func PromptBool(label string, defaultValue bool) bool {
	cyan := color.New(color.FgCyan).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	scanner := bufio.NewScanner(os.Stdin)
	defaultStr := "n"
	if defaultValue {
		defaultStr = "y"
	}
	fmt.Printf("%s (y/n) [%s]: ", cyan(label), yellow(defaultStr))

	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())

	if input == "" {
		return defaultValue
	}
	input = strings.ToLower(input)
	return input == "y" || input == "yes"
}
