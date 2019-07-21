package ftp

import (
	"bufio"
	"os"
)

type GetCommand struct {
	Command
	FileMeta
	path        string
	bytesCopied int64
}

func (gc *GetCommand) Name() string {
	return "Get Command"
}

func (gc *GetCommand) String() string {
	return ""
}

func (gc *GetCommand) Execute() error {
	file, err := os.Open(gc.path)
	if err != nil {
		return err
	}
	if err := gc.setFile(file); err != nil {
		return err
	}

	if err := gc.sendFile(); err != nil {
		return err
	}

	return nil
}

func (gc *GetCommand) setFile(file *os.File) error {
	gc.file = file
	fileStat, err := file.Stat()
	if err != nil {
		return err
	}
	gc.fileSize = fileStat.Size()

	return nil
}

func (gc *GetCommand) sendFile() error {
	r := bufio.NewReader(gc.file)
	w := bufio.NewWriter(gc.session.conn)

	buf := make([]byte, 4096)

	for {
		_, err := r.Read(buf)
		if err != nil {
			return err
		}

		n, err := w.Write(buf)
		if err != nil {
			return err
		}

		gc.bytesCopied += int64(n)
		buf = buf[:0]
	}
}
