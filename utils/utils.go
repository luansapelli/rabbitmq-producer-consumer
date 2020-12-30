package utils

import "log"

func HandleError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", err, msg)
	}
}
