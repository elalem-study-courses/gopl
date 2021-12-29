package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type rootInfo struct {
	dir    string
	nbytes int64
	nfiles int64
}

func walkDir(dir string, fileSizes chan<- rootInfo, rootInfoMap map[string]*rootInfo, key string) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes, rootInfoMap, key)
		} else {
			info := rootInfo{dir: key, nbytes: entry.Size()}
			fileSizes <- info
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}

	return entries
}

func printDiskUsage(info *rootInfo) {
	fmt.Printf("%d files, %.1f GB\n", info.nfiles, float64(info.nbytes)/1e9)
}

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan rootInfo)
	rootInfoMap := make(map[string]*rootInfo)
	go func() {
		for _, root := range roots {
			rootInfoMap[root] = &rootInfo{dir: root, nbytes: 0, nfiles: 0}
			walkDir(root, fileSizes, rootInfoMap, root)
		}

		close(fileSizes)
	}()

loop:
	for {
		select {
		case info, ok := <-fileSizes:
			if !ok {
				break loop
			}

			// rInfo := rootInfoMap
			rootInfoMap[info.dir].nbytes += info.nbytes
			rootInfoMap[info.dir].nfiles++
		case <-tick:
			for _, root := range roots {
				fmt.Printf("Info for %s\n", root)
				printDiskUsage(rootInfoMap[root])
			}
		}
	}

	for _, root := range roots {
		fmt.Printf("Info for %s\n", root)
		printDiskUsage(rootInfoMap[root])
	}
}
