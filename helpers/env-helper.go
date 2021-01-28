package helpers

import "os"

func SetEnv() {
	os.Setenv("EMAIL_PASSWORD", "darrowred29")
	os.Setenv("IMAGE_PATH", "C:/DATA/Golang/src/github.com/lionelritchie29/staem-backend/images")
}

