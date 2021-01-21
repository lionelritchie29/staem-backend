package helpers

import "os"

func SetEnv() {
	os.Setenv("MAILJET_PUBLIC_API_KEY", "988e2aa60ce2b596f70a93c7af6e46fc")
	os.Setenv("MAILJET_PRIVATE_API_KEY", "f47d062762e099dd204e63e871acd14b")
	os.Setenv("IMAGE_PATH", "C:/DATA/Golang/src/github.com/lionelritchie29/staem-backend/images")
}

