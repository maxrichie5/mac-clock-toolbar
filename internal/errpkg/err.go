package errpkg

import "log"

func CheckFatalErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
