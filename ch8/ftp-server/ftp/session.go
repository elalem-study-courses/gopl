package ftp

import (
	"fmt"
	"net"
	"os"
	"time"
)

var (
	registeredSessions map[int64]*Session
)

func init() {
	registeredSessions = make(map[int64]*Session)
}

type Commands []Executor

// func (c Commands) String() string {
// 	var buf bytes.Buffer
// 	for _, command := range c {
// 		fmt.Fprintf(&buf, "%v\n", command)
// 	}
// 	return buf.String()
// }

type Session struct {
	conn      net.Conn
	id        int64
	createdAt time.Time
	Working   Commands
	path      string
}

func (s *Session) Start() error {
	registerSession(s)
	msg := fmt.Sprintf("Started session with %v successfully", s.conn.RemoteAddr())
	fmt.Fprint(s.conn, msg)
	fmt.Fprintln(os.Stdout, msg)

	return nil
}

func NewSession(conn net.Conn) *Session {
	return &Session{
		conn:      conn,
		id:        time.Now().Unix(),
		createdAt: time.Now(),
		path:      FTPServerRoot,
	}
}

func (s *Session) Close() error {
	// unregisterSession(s)
	fmt.Fprintf(os.Stdout, "Closing connection with %v\n", s.conn.RemoteAddr())
	return s.conn.Close()
}

func (s *Session) Handle(line string) error {
	fmt.Printf("Received command %q from %v\n", string(line), s.conn.RemoteAddr())
	command, err := newCommand(s, line)
	if err != nil {
		return fmt.Errorf("Error while creating command %q: %v", line, err)
	}

	s.registerCommnad(command)

	if err := command.Execute(); err != nil {
		return fmt.Errorf("Error while executing command %q: %v", line, err)
	}
	if err := s.unregisterCommand(command); err != nil {
		return err
	}
	return nil
}

func (s *Session) registerCommnad(cmd Executor) {
	s.Working = append(s.Working, cmd)
}

func (s *Session) unregisterCommand(cmd Executor) error {
	s.Working = rejectElements(s.Working, func(a Executor) bool {
		return a.ID() == cmd.ID()
	})
	return nil
}

func (s *Session) writeString(str string) {
	s.conn.Write([]byte(str))
}

func registerSession(s *Session) {
	registeredSessions[s.id] = s
}

func unregisterSession(s *Session) {
	delete(registeredSessions, s.id)
}

func rejectElements(commands Commands, cmp func(Executor) bool) Commands {
	selectedCommands := make(Commands, 0)
	for _, command := range commands {
		if !cmp(command) {
			selectedCommands = append(selectedCommands, command)
		}
	}

	return selectedCommands
}
