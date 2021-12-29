package server

import (
	"fmt"
	"net"
	"os"
	"strings"
	"text/tabwriter"
)

var lsCommandColumn = []string{
	"IsDirectory",
	"Name",
}

type ListCommand struct{}

func (lc *ListCommand) handle(conn net.Conn) {
	entries, err := os.ReadDir(DirectoryRoot)
	if err != nil {
		fmt.Fprintln(conn, err)
	}
	writer := tabwriter.NewWriter(conn, 0, 4, 4, ' ', tabwriter.TabIndent)

	fmt.Fprintln(writer, strings.Join(lsCommandColumn, "\t"))
	for _, entry := range entries {
		isDirectory := false
		if entry.IsDir() {
			isDirectory = true
		}

		fmt.Fprintf(writer, "%t\t%s\n", isDirectory, entry.Name())
	}

	writer.Flush()
}

func init() {
	commands["ls"] = &ListCommand{}
}
