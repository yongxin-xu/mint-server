package mintcommon

import (
	"fmt"
	"sync"
	"time"
)

// Mutex locks print functions
var Mutex sync.Mutex

// DebugPrint logs server's information to console
// it prints log to console or file based on user's choice
func DebugPrint(doLog bool, toConsole bool, path string, a string) {
	if doLog {
		if toConsole {
			logToConsole(a)
		} else {
			logToFile(path, a)
		}
	}
}

// logToConsole logs server's information to console
func logToConsole(a string) {
	current_time := time.Now()
	year := current_time.Year()
	month := current_time.Month()
	day := current_time.Day()
	hour := current_time.Hour()
	minute := current_time.Minute()
	second := current_time.Second()
	Mutex.Lock()
	fmt.Printf("%d-%s-%d %d:%d:%d    %s\n", year, month, day, hour, minute, second, a)
	Mutex.Unlock()
}

// logToFile logs server's information to file
func logToFile(path string, a ...interface{}) {
	// TODO: implement logtofile
}
