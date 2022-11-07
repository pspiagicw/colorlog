package colorlog

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)



func LogSuccess(format string, v ...interface{}) {
    message := fmt.Sprintf(format, v...)
    color.Green(message)
}

func LogError(format string, v ...interface{}) {
    message := fmt.Sprintf(format , v...)

    color.Red(message)
}
func LogFatal(format string, v ...interface{}) {
    message := fmt.Sprintf(format , v...)
    color.Red(message)

    os.Exit(1)
}
func LogInfo(format string, v ...interface{}) {
    message := fmt.Sprintf(format , v...)
    color.Yellow(message)
}
func LogSuccessFancy(format string, v ...interface{}) {
	message := "[*] ✓ " + fmt.Sprintf(format, v...)
	color.Green(message)
}
func LogErrorFancy(format string, v ...interface{}) {
	message := "[*] x " + fmt.Sprintf(format, v...)
	color.Red(message)
}
func LogInfoFancy(format string, v ...interface{}) {
	message := "[*] + " + fmt.Sprintf(format, v...)
	color.Yellow(message)
}
