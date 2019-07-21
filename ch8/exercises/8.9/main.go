package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sync"
	"time"

	"os"
)

type DirectoryInfo struct {
	nfiles int64
	nbytes int64
}

type RootDirectoryFileInfo struct {
	root     string
	fileSize int64
}

var DirInfoMap map[string]DirectoryInfo

func walkDir(root string, wg *sync.WaitGroup, fileSizes chan<- RootDirectoryFileInfo, mainRoot string) {
	defer wg.Done()

	for _, entry := range dirents(root) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(root, entry.Name())
			go walkDir(subdir, wg, fileSizes, mainRoot)
		} else {
			fileSizes <- RootDirectoryFileInfo{root: mainRoot, fileSize: entry.Size()}
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(root string) []os.FileInfo {
	sema <- struct{}{}
	defer func() {
		<-sema
	}()

	entries, err := ioutil.ReadDir(root)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du4: %v\n", err)
	}

	return entries
}

var verbose = flag.Bool("v", false, "show verbose progress messages")
var highVerbose = flag.Bool("vv", false, "display sizes for each root")

func main() {
	flag.Parse()
	roots := flag.Args()

	if len(roots) == 0 {
		roots = []string{"."}
	}

	DirInfoMap = map[string]DirectoryInfo{}

	fileSizes := make(chan RootDirectoryFileInfo)

	var wg sync.WaitGroup

	var tick <-chan time.Time

	if *verbose || *highVerbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	for _, root := range roots {
		wg.Add(1)
		go func(root string) {
			walkDir(root, &wg, fileSizes, root)
		}(root)
	}

	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	var nfiles, nbytes int64

loop:
	for {
		select {
		case rootInfo, ok := <-fileSizes:
			if !ok {
				break loop
			}

			nfiles++
			nbytes += rootInfo.fileSize

			DirInfoMap[rootInfo.root] = DirectoryInfo{nbytes: nbytes, nfiles: nfiles}

		case <-tick:
			if *highVerbose {
				for _, root := range roots {
					printDiskUsageForDirectory(root, DirInfoMap[root])
				}
			} else {
				printDiskUsage(nfiles, nbytes)
			}
		}
	}

	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Fprintf(os.Stdout, "%d files\t%.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func printDiskUsageForDirectory(root string, directoryInfo DirectoryInfo) {
	fmt.Fprintf(os.Stdout, "%s\t%d files\t%.1f GB\n", root, directoryInfo.nfiles, float64(directoryInfo.nbytes)/1e9)
}
