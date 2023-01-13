package statikFs

import "net/http"

type DirHttpFs struct {
	path string
}

func (d *DirHttpFs) Open(name string) (http.File, error) {
	return statikFS.Open(d.path + name)
}
