package target

import "github.com/jakehl/goid"

func AwesomeUUID() string {
	return goid.NewV4UUID().String()
}
