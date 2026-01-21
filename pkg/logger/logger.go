package logger

import (
	"fmt"
	"log"
	"os"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorCyan   = "\033[36m"
	colorGray   = "\033[90m"
)

func init() {
	// Remove date/time from logs for cleaner output if desired, 
	// but let's keep it simple and just add color.
	log.SetOutput(os.Stderr)
}

func Success(format string, v ...interface{}) {
	log.Printf(colorGreen+"[SUCCESS] "+format+colorReset, v...)
}

func Info(format string, v ...interface{}) {
	log.Printf(colorBlue+"[INFO] "+format+colorReset, v...)
}

func Warn(format string, v ...interface{}) {
	log.Printf(colorYellow+"[WARN] "+format+colorReset, v...)
}

func Error(format string, v ...interface{}) {
	log.Printf(colorRed+"[ERROR] "+format+colorReset, v...)
}

func Debug(format string, v ...interface{}) {
	log.Printf(colorGray+"[DEBUG] "+format+colorReset, v...)
}

func Phase(format string, v ...interface{}) {
	log.Printf(colorCyan+colorReset+colorCyan+"[PHASE] "+format+colorReset, v...)
}

func Header(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	log.Printf(colorCyan + "========================================" + colorReset)
	log.Printf(colorCyan + "   " + msg + colorReset)
	log.Printf(colorCyan + "========================================" + colorReset)
}
