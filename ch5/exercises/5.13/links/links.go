package links

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"

	"net/url"

	"golang.org/x/net/html"
)

var savingDirectory string

// Sets the directory where the content will be saved
func SetSavingDirectory(dir string) {
	savingDirectory = dir
}

// Extract makes an HTTP GET request to the specified url...
func Extract(URL string, takeSnapshot bool) ([]string, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", URL, resp.Status)
	}

	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	doc, err := html.Parse(bytes.NewReader(responseBytes))

	if takeSnapshot {
		if err := snapshotResponse(URL, bytes.NewReader(responseBytes)); err != nil {
			return nil, fmt.Errorf("Couldn't snapshot response for %s: %v", URL, err)
		}

	}

	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", URL, err)
	}

	// This array will be visible to visitNode wherever it was used.
	// So it will be populated with the links even though the forEach function
	// does not return the links
	var links []string

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(attr.Val)
				if err != nil {
					continue // Bad url
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(*html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func snapshotResponse(URL string, reader io.Reader) error {
	directoryExists := isFileSystemExists(savingDirectory)
	if !directoryExists {
		err := createDirectory(savingDirectory)
		if err != nil {
			return fmt.Errorf("Couldn't take snapshot of %s: %v", URL, err)
		}
	}

	host := parseURL(URL)
	snapshotDirectory := path.Join(savingDirectory, host)
	snapshotDirectoryExists := isFileSystemExists(snapshotDirectory)
	if !snapshotDirectoryExists {
		err := createDirectory("./" + snapshotDirectory)
		if err != nil {
			return fmt.Errorf("Couldn't create snapshot directory %s: %v", snapshotDirectory, err)
		}
	}

	filename := strings.ReplaceAll(URL, "/", "_")
	if err := saveFile(snapshotDirectory, filename, reader); err != nil {
		return fmt.Errorf("Couldn't save file in %s: %v", snapshotDirectory, err)
	}
	return nil
}

func isFileSystemExists(fileSystem string) bool {
	_, err := os.Stat(fileSystem)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func createDirectory(dir string) error {
	if err := os.Mkdir(dir, 0755); err != nil {
		return fmt.Errorf("Couldn't create directory %s: %v", dir, err)
	}
	return nil
}

func parseURL(URL string) string {
	u, _ := url.Parse(URL)
	return u.Host
}

func saveFile(dir, filename string, reader io.Reader) error {
	filePath := path.Clean(path.Join(dir, filename))
	if fileExists := isFileSystemExists(filePath); fileExists {
		return fmt.Errorf("Duplicated %s", filePath)
	}

	file, err := createFile(filePath)
	if err != nil {
		return fmt.Errorf("Couldn't save file %s: %v", filePath, err)
	}
	if _, err := io.Copy(file, reader); err != nil {
		return fmt.Errorf("Couldn't write to file %s: %v", filePath, err)
	}

	file.Close()
	return nil
}

func createFile(filename string) (*os.File, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, fmt.Errorf("Couldn't create file %s: %v", filename, err)
	}
	return file, nil
}
