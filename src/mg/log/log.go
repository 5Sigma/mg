package console

import (
	"fmt"
	"github.com/ttacon/chalk"
	"os"
)

func Print(msg string, args ...interface{}) {
	fmt.Println(msg)
}

func Printf(msg string, args ...interface{}) {
	Print(fmt.Sprintf(msg, args...))
}

func Warn(msg string) {
	Print(chalk.Yellow.Color(msg))
}

func Warnf(msg string, args ...interface{}) {
	Warn(fmt.Sprintf(msg, args...))
}

func Alert(msg string) {
	Print(chalk.Bold.TextStyle(msg))
}

func Alertf(msg string, args ...interface{}) {
	Alert(fmt.Sprintf(msg, args...))
}

func Error(msg string) {
	Print(chalk.Red.Color(msg))
}

func Errorf(msg string, args ...interface{}) {
	Error(fmt.Sprintf(msg, args...))
}

func Success(msg string) {
	Print(chalk.Green.Color(msg))
}

func Successf(msg string, args ...interface{}) {
	Success(fmt.Sprintf(msg, args...))
}
func Fatal(msg string) {
	Print(chalk.Red.Color(msg))
	os.Exit(-1)
}

func Fatalf(msg string, args ...interface{}) {
	Fatal(fmt.Sprintf(msg, args...))
}
