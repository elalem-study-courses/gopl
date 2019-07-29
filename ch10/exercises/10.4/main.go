package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"unicode"
	"unicode/utf8"
)

/**
Two approaches can be followed
1. Just use Deps entry in go list which already contains all the dependencies of a package
2. Use DFS to traverse imports (will follow this approach)
*/

type Package struct {
	ImportPath string
	Imports    []string
}

type Packages []*Package

/**
I did not initialize a visited map because the tool is supposed to print
all the dependencies of each requested package
so it will be an empty map for each run
*/
var (
	graph map[string][]string
	used  map[string]map[string]struct{}
)

func init() {
	graph = make(map[string][]string)
	used = make(map[string]map[string]struct{})
}

func main() {
	packages, err := loadAllPackages()
	if err != nil {
		log.Fatal(err)
	}

	populateDependencyGraph(packages)

	requestedPackages, err := loadRequestedPackages()

	for _, requestedPackage := range requestedPackages {
		traverse(requestedPackage, make(map[string]struct{}))
		fmt.Printf("--------------------------------------\n\n\n")
	}
}

func loadAllPackages() (Packages, error) {
	allPackagesCommand := exec.Command("go", "list", "-json", "...")
	out, err := allPackagesCommand.Output()
	if err != nil {
		return nil, fmt.Errorf("error in command: %v", err)
	}
	innerJSONObjects := strings.Join(strings.Split(removeAllWhiteSpaces(string(out)), "]}"), "]},")
	correctJSON := "[" + innerJSONObjects[:len(innerJSONObjects)-1] + "]"

	packages := make(Packages, 0)

	if err := json.NewDecoder(strings.NewReader(correctJSON)).Decode(&packages); err != nil {
		return nil, fmt.Errorf("error in unmarshaling: %v", err)
	}

	return packages, nil
}

func addUsedPair(from, to string) {
	if used[from] == nil {
		used[from] = make(map[string]struct{})
	}
	used[from][to] = struct{}{}
}

func populateDependencyGraph(packages Packages) {
	for _, currentPackage := range packages {
		for _, importedPackageName := range currentPackage.Imports {
			if _, ok := used[currentPackage.ImportPath][importedPackageName]; !ok {
				addUsedPair(currentPackage.ImportPath, importedPackageName)
				graph[currentPackage.ImportPath] = append(graph[currentPackage.ImportPath], importedPackageName)
			}
		}
	}

}

func removeAllWhiteSpaces(s string) string {
	var out bytes.Buffer
	for _, ch := range s {
		if utf8.ValidRune(ch) && !unicode.IsSpace(ch) && ch != '\n' {
			out.WriteRune((ch))
		}
	}

	return out.String()
}

func traverse(cur string, visited map[string]struct{}) {
	if _, ok := visited[cur]; !ok {
		visited[cur] = struct{}{}
		fmt.Println(cur)
	}

	for _, child := range graph[cur] {
		traverse(child, visited)
	}
}

func loadRequestedPackages() ([]string, error) {
	args := []string{"list"}
	args = append(args, os.Args[1:]...)
	cmd := exec.Command("go", args...)
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	packages := strings.Split(strings.TrimSpace(string(out)), "\n")

	return packages, nil
}
