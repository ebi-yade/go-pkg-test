package ex

import "github.com/uniplaces/carbon"

func CarbonTime() string {
	return carbon.Now().DateTimeString()
}
