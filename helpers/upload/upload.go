package upload

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

// Remove - remove a file from a directory
func Remove(filename string, directory string) error {
	err := os.Remove(fmt.Sprintf("./%s/%s", directory, filename))

	return err
}

// Upload - create and upload a file
func Upload(file multipart.File, avatarFile string, directory string) error {
	defer file.Close()

	out, err := os.Create(fmt.Sprintf("./%s/%s", directory, avatarFile))
	if err != nil {
		return err
	}
	defer out.Close()

	// write the content from POST to the file
	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}

	return nil
}
