package zip

import (
	"archive/zip"
	"fmt"

	archiveReader "learning/ch8/exercises/10.2/archive-reader"
)

func init() {
	archiveReader.RegisterReader("zip", reader)
}

func reader(filePath string) error {
	zipArchive, err := zip.OpenReader(filePath)
	if err != nil {
		return err
	}
	defer zipArchive.Close()

	for _, f := range zipArchive.File {
		fmt.Println(f.Name)
	}
	return nil
}
