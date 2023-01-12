package statikFs

import (
	"fmt"
	"io/ioutil"
	"juejinCollections/tool"
	"net/http"

	"github.com/rakyll/statik/fs"

	_ "juejinCollections/server/statikFs/statik"
)

var statikFS http.FileSystem

type DirHttpFs struct {
	path string
}

func (d *DirHttpFs) Open(name string) (http.File, error) {
	return statikFS.Open(d.path + name)
}

func InitStatikFs() {
	s, err := fs.New()
	if err != nil {
		tool.PanicErr(err)
	}
	statikFS = s
}

func OpenDir(name string) http.FileSystem {
	return &DirHttpFs{
		path: name,
	}
}

func GetFileSystem() http.FileSystem {
	return statikFS
}

func GetFile() {
	r, err := statikFS.Open("/index.html")
	if err != nil {
		tool.PanicErr(err)
	}
	defer r.Close()
	contents, err := ioutil.ReadAll(r)
	if err != nil {
		tool.PanicErr(err)
	}

	fmt.Println(string(contents))
}
