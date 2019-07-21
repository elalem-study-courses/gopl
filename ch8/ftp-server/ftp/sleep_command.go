package ftp

import (
	"fmt"
	"time"
)

type SleepCommand struct {
	Command
	duration time.Duration
	session  *Session
}

func (sc *SleepCommand) Execute() error {
	time.Sleep(sc.duration)
	return nil
}

func (sc *SleepCommand) String() string {
	return fmt.Sprintf("Sleeping for %v, Elapsed time %v", sc.duration, time.Since(sc.createdAt))
}

func (sc *SleepCommand) Name() string {
	return "Sleep Command"
}

func (sc *SleepCommand) ID() int64 {
	return sc.id
}
