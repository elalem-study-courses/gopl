package ftp

import (
	"bytes"
	"fmt"
	"text/tabwriter"
	"time"
)

type History struct {
	command  string
	issuedAt time.Time
}

func (h *History) String() string {
	var bf *bytes.Buffer
	tw := tabwriter.NewWriter(bf, 2, 2, 2, ' ', tabwriter.TabIndent)
	fmt.Fprintf(tw, ">\t%s\t%v", h.command, time.Since(h.issuedAt))
	tw.Flush()
	return bf.String()
}
