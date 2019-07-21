package ftp

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Namer interface {
	Name() string
}

type Executor interface {
	Execute() error
	String() string
	ID() int64
	Namer
}

type UnknownCommand struct {
	command string
}

func (uc *UnknownCommand) Error() string {
	return fmt.Sprintf("Unknown command %v", uc.command)
}

type Command struct {
	arguments []string
	createdAt time.Time
	session   *Session
	id        int64
}

func (c *Command) String() string { return "" }
func (c *Command) Name() string   { return "" }
func (c *Command) Execute() error { return nil }

func newCommand(session *Session, cmd string) (Executor, error) {
	var physicalCommand Executor

	tokens := strings.Split(cmd, " ")

	metaCommand := Command{
		createdAt: time.Now(),
		session:   session,
		id:        time.Now().Unix(),
	}
	switch tokens[0] {
	case "ls":
		metaCommand.arguments = tokens[1:]
		physicalCommand = &ListCommand{
			Command: metaCommand,
		}
	case "sleep":
		if len(tokens) < 2 {
			return nil, fmt.Errorf("Expected 2 arguments found %d", len(tokens))
		}
		seconds, err := strconv.Atoi(tokens[1])
		if err != nil {
			return nil, err
		}
		physicalCommand = &SleepCommand{
			duration: time.Duration(seconds) * time.Second,
			Command:  metaCommand,
		}
	case "status":
		physicalCommand = &StatusCommand{
			Command: metaCommand,
		}

	default:
		return nil, &UnknownCommand{command: tokens[0]}
	}

	return physicalCommand, nil
}
