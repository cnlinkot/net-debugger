package util

import "log"

func CheckFatalError(err error, message ...any) {
	if err != nil {
		log.Fatal(message, ": ", err.Error())
	}
}
func CheckError(err error, message ...any) {
	if err != nil {
		log.Println(message, ": ", err.Error())
	}
}
