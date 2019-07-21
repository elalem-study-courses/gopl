package ftp

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

type ListCommand struct {
	Command
}

func (ls *ListCommand) Execute() error {
	var buf bytes.Buffer

	path := ls.session.path

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	fmt.Fprintf(&buf, "Listing files in %s\n", ls.session.path)

	for _, file := range files {
		isDir := '-'
		if file.IsDir() {
			isDir = 'd'
		}
		fmt.Fprintf(&buf, "%c\t%s\n", isDir, file.Name())
	}

	ls.session.writeString(buf.String())
	return nil
}

func (ls *ListCommand) Name() string {
	return "List Command"
}

func (ls *ListCommand) String() string {
	return fmt.Sprintf("Listing directory %s", ls.session.path)
}

func (ls *ListCommand) ID() int64 {
	return ls.id
}
