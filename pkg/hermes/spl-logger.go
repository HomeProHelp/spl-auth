package hermes

import "log"

var logLevel [3]string = [3]string{"INFO", "WARN", "ERROR"}

func Log(level int, message string, fatal bool) {
	if fatal {
		log.Fatalf("\nLEVEL=%s\nMESSAGE=%s\n\n", logLevel[level-1], message)
		return
	}
	log.Printf("\nLEVEL=%s\nMESSAGE=%s\n\n", logLevel[level-1], message)

}
