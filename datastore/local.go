package datastore

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/k1LoW/octocov/report"
)

type Local struct {
	root string
}

func NewLocal(root string) (*Local, error) {
	fi, err := os.Stat(root)
	if err != nil {
		return nil, err
	}
	if !fi.IsDir() {
		return nil, fmt.Errorf("%s is not directory", root)
	}
	return &Local{
		root: root,
	}, nil
}

func (l *Local) Store(ctx context.Context, path string, r *report.Report) error {
	return os.WriteFile(filepath.Join(l.root, path), r.Bytes(), os.ModePerm)
}

func (l *Local) FS() (fs.FS, error) {
	return &LocalFS{
		root: filepath.Join(l.root),
	}, nil
}

type LocalFS struct {
	root string
}

func (fsys *LocalFS) Open(name string) (fs.File, error) {
	f, err := os.Open(filepath.Clean(filepath.Join(fsys.root, name)))
	if f == nil {
		return nil, err
	}
	return f, err
}
