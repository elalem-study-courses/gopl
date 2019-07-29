package tar

import (
	tarpkg "archive/tar"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	archiveReader "learning/ch8/exercises/10.2/archive-reader"
)

func init() {
	archiveReader.RegisterReader("tar", reader)
}

func reader(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	tarArchive := tarpkg.NewReader(file)
	for {
		header, err := tarArchive.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}
		fmt.Println(header.Name)
		io.Copy(ioutil.Discard, tarArchive)
	}

	return nil
}
