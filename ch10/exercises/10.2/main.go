package main

import (
	"flag"
	"fmt"

	_ "learning/ch8/exercises/10.2/archive-reader/tar"
	_ "learning/ch8/exercises/10.2/archive-reader/zip"

	reader "learning/ch8/exercises/10.2/archive-reader"
)

/**
This package was put in go path using symbolic links to fix relative import issue.
*/

var fileName = flag.String("file-name", "", "file to read")

func main() {
	flag.Parse()

	errs := reader.TryAllForFile(*fileName)
	fmt.Println(errs)
}
