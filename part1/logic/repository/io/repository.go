package io

import (
	"io/fs"
	"os"
)

type Repository interface {
	ReadDir(dirname string) ([]fs.FileInfo, error)
	ReadAll(file *os.File) ([]byte, error)

	OsClose(file *os.File) error
	OsOpen(name string) (*os.File, error)
}
