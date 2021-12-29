package mirrors

import (
	"io/ioutil"
	"net/url"
	"os"
	"strings"
)

func generatePath(link *url.URL) string {
	folderPath := link.String()
	folderPath = strings.TrimPrefix(folderPath, "https://")
	folderPath = strings.TrimPrefix(folderPath, "http://")

	return folderPath
}

func createFolderForLink(folderPath string) {
	os.MkdirAll(folderPath, os.ModePerm)
}

func writeFile(filePath string, content []byte) {
	ioutil.WriteFile(filePath, content, os.ModePerm)
}
