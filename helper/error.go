package helper

import "log"

func FailOnError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", err, msg)
	}
}
