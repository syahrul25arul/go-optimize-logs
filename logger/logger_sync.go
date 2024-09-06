package logger

// import (
// 	"fmt"
// 	"log"
// )

// const (
// 	Black   = "\u001b[97;100m"
// 	Red     = "\u001b[37;101m"
// 	Green   = "\u001b[37;102m"
// 	Yellow  = "\u001b[37;103m"
// 	Blue    = "\u001b[37;44m"
// 	Magenta = "\u001b[97;105m"
// 	Cyan    = "\u001b[97;106m"
// 	White   = "\u001b[97m"
// 	Reset   = "\u001b[0m"
// )

// type LogType func() string

// func LogInfo() string {
// 	return fmt.Sprintf("|%s INFO %s|", Blue, Reset)
// }

// func LogError() string {
// 	return fmt.Sprintf("|%s ERROR %s|", Red, Reset)
// }

// func LogPanic() string {
// 	return fmt.Sprintf("|%s PANIC %s|", Red, Reset)
// }

// func LogWarning() string {
// 	return fmt.Sprintf("|%s Warning %s|", Yellow, Reset)
// }
// func WriteLog(msg string, logType LogType, args string) {
// 	log.Printf("%s %s - %s", logType(), msg, args)
// }
