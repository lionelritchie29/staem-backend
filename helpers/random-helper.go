package helpers

import "math/rand"

func RandBool() bool {
	rInt := rand.Intn(2)
	if rInt == 1 {
		return  true
	}
	return false
}
