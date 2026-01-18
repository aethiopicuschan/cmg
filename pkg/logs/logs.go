package logs

import (
	"fmt"
	"os"
	"time"

	"github.com/aethiopicuschan/cmg/pkg/internal/colors"
)

func Info(msg string) {
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s %s %s\n", colors.Blue("[INFO]"), colors.Cyan(now), colors.Green(msg))
}

func Warn(msg string) {
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s %s %s\n", colors.Yellow("[WARN]"), colors.Cyan(now), colors.Yellow(msg))
}

func Error(msg string) {
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s %s %s\n", colors.Red("[ERROR]"), colors.Cyan(now), colors.Yellow(msg))
}

func Fatal(msg string) {
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s %s %s\n", colors.Red("[FATAL]"), colors.Cyan(now), colors.Yellow(msg))
	os.Exit(1)
}
