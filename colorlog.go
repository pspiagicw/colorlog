package colorlog

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

// Log when something goes right
// Prints using green color
func LogSuccess(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	color.Green(message)
}

// Log when something goes wrong
// Prints using a red color
func LogError(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)

	color.Red(message)
}

// Log when something does very wrong
// Prints using red color
// Additionally exits from the application using `os.Exit`
func LogFatal(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	color.Red(message)

	os.Exit(1)
}

// Log when communicating some info
// Prints using a yellow color
func LogInfo(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	color.Yellow(message)
}

// Log something positive
// Prints in green with a ASCII tick (`✓`)
func LogSuccessFancy(format string, v ...interface{}) {
	message := "[*] ✓ " + fmt.Sprintf(format, v...)
	color.Green(message)
}

// Log something negative
// Prints in red with a ASCII X (`x`)
func LogErrorFancy(format string, v ...interface{}) {
	message := "[*] x " + fmt.Sprintf(format, v...)
	color.Red(message)
}

// Log something informative
// Prints in yello in ASCII plug (`+`)
func LogInfoFancy(format string, v ...interface{}) {
	message := "[*] + " + fmt.Sprintf(format, v...)
	color.Yellow(message)
}

// ColorLogger interface for dependency injection purposes.
// Useful for mocking with custom type.
type ColorLogger interface {
    LogSuccess(string, ...interface{})
    LogError(string, ...interface{})
    LogInfo(string, ...interface{})
    LogFatal(string, ...interface{})

    LogSuccessFancy(string, ...interface{})
    LogErrorFancy(string, ...interface{})
    LogInfoFancy(string, ...interface{})
}


// Empty struct for implementing the ColorLogger interface.
type DefaultColorLogger struct {
}

func (d *DefaultColorLogger) LogSuccess(format string, v ...interface{}) {
    LogSuccess(format , v...)
}
func (d *DefaultColorLogger) LogError(format string, v ...interface{}) {
    LogError(format , v...)
}
func (d *DefaultColorLogger) LogInfo(format string, v ...interface{}) {
    LogInfo(format, v...)
}
func (d *DefaultColorLogger) LogFatal(format string, v ...interface{}) {
    LogFatal(format, v...)
}



func (d *DefaultColorLogger) LogSuccessFancy(format string, v ...interface{}) {
    LogSuccessFancy(format, v...)
}
func (d *DefaultColorLogger) LogErrorFancy(format string, v ...interface{}) {
    LogErrorFancy(format, v...)
}
func (d *DefaultColorLogger) LogInfoFancy(format string, v ...interface{}) {
    LogInfoFancy(format, v...)
}
