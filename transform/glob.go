package transform

import "path/filepath"

type Globber interface {
    Glob(pattern string) (matches []string, err error)
}

func NewGlobber() Globber {
    return &pathGlobber{}
}

type pathGlobber struct{}

func (pathGlobber) Glob(pattern string) (matches []string, err error) {
    return filepath.Glob(pattern)
}
