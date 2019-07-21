package ftp

import "os"

type FileMeta struct {
	file     *os.File
	fileSize int64
}
