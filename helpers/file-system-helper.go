package helpers

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func MakeDirectoryIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, os.ModeDir|0755)
	}
	return nil
}

func SaveImage(encodedBase64Img string, filename string) {
	decoded, _ := base64.StdEncoding.DecodeString(encodedBase64Img)

	writeFilename := os.Getenv("IMAGE_PATH") + "/" + filename
	fmt.Println(writeFilename)
	errs := ioutil.WriteFile( writeFilename, decoded, os.ModePerm)

	if errs != nil {
		fmt.Print("Error Save File: ")
		fmt.Println(errs)
	}

}
