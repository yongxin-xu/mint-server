package mintcommon

import (
	"fmt"
	"sync"
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
	Mutex.Lock()
	fmt.Println(a)
	Mutex.Unlock()
}

// logToFile logs server's information to file
func logToFile(path string, a ...interface{}) {
	// TODO: implement logtofile
}
