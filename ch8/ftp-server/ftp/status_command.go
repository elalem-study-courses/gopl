package ftp

import (
	"bytes"
	"fmt"
	"text/tabwriter"
	"time"
)

type StatusCommand struct {
	Command
}

func (sc *StatusCommand) Execute() error {
	var buf bytes.Buffer
	tw := tabwriter.NewWriter(&buf, 4, 4, 4, ' ', 0)
	fmt.Fprintf(tw, "\tSession is working for %v\n", time.Since(sc.session.createdAt))

	fmt.Fprintf(tw, "\tCommand Name\tCommand Status\t\n")
	for _, command := range sc.session.Working {
		fmt.Fprintf(tw, "\t%v\t%v\t\n", command.Name(), command)
	}
	tw.Flush()
	sc.session.writeString(buf.String())
	return nil
}

func (sc *StatusCommand) String() string {
	return ""
}

func (sc *StatusCommand) Name() string {
	return "Status Commnad"
}

func (sc *StatusCommand) ID() int64 {
	return sc.id
}
