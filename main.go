package mylogger

import "log"

func LogInfo(messsage string) {
	log.Printf("INFO = %v", messsage)
}
