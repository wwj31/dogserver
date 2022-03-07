package controller

import "fmt"

func assert(b bool, log ...string) {
	if !b {
		str := "assert failed "
		if len(log) > 1 {
			str = fmt.Sprintf(log[0], log[1:])
		} else if len(log) > 0 {
			str = log[0]
		}
		panic(str)
	}
}
