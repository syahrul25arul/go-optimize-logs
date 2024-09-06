package logger

import (
	"fmt"
	"log"
	"sync"
)

const (
	Black   = "\u001b[97;100m"
	Red     = "\u001b[37;101m"
	Green   = "\u001b[37;102m"
	Yellow  = "\u001b[37;103m"
	Blue    = "\u001b[37;44m"
	Magenta = "\u001b[97;105m"
	Cyan    = "\u001b[97;106m"
	White   = "\u001b[97m"
	Reset   = "\u001b[0m"
)

type LogType func() string

func LogInfo() string {
	return fmt.Sprintf("|%s INFO %s|", Blue, Reset)
}

func LogError() string {
	return fmt.Sprintf("|%s ERROR %s|", Red, Reset)
}

func LogPanic() string {
	return fmt.Sprintf("|%s PANIC %s|", Red, Reset)
}

func LogWarning() string {
	return fmt.Sprintf("|%s Warning %s|", Yellow, Reset)
}
func (l *Log) WriteLog(msg string, logType LogType, args string) {
	l.wg.Add(1)
	go func() {
		// for push data to channel, don't use struct pointer.
		// use struct or better is primitive type
		defer l.wg.Done()
		l.logMsg <- fmt.Sprintf("%s %s - %s", logType(), msg, args)
	}()

}

type Log struct {
	logMsg chan interface{}
	// lc     *sync.Locker
	wg *sync.WaitGroup
}

func InitLog() *Log {
	log := &Log{
		wg:     &sync.WaitGroup{},
		logMsg: make(chan interface{}),
	}
	go log.LogRuntime()
	return log
}

func (l *Log) LogRuntime() {
	for data := range l.logMsg {
		if data == nil {
			close(l.logMsg)
			return
		}
		if msg, ok := data.(string); ok {
			log.Print(msg)
		}
	}
}

func (l *Log) CloseLog() {
	l.wg.Wait()
	l.logMsg <- nil
}
