package shared

import "log"

func RecoverApp() {
	if r := recover(); r != nil {
		log.Println("error: ", r)
	}
}