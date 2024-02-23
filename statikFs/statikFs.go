package statikFs

import (
	"io"
	"os"
	"juejinCollections/config"
	"juejinCollections/tool"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rakyll/statik/fs"

	_ "juejinCollections/statikFs/statik"
)

var statikFS http.FileSystem

func InitStatikFs() {
	s, err := fs.New()
	if err != nil {
		tool.PanicErr(err)
	}
	statikFS = s
}

// 在statik中一切文件都已 / 开头
func GetStatikPath(path string) string {
	return "/" + tool.RegExpReplace("^(./|../)", path, "")
}

func OpenDir(name string) http.FileSystem {
	return &DirHttpFs{
		path: GetStatikPath(name),
	}
}

func GetFileSystem() http.FileSystem {
	return statikFS
}

func GetFileData(file string) ([]byte, error) {
	if config.Config.IsDevelopment {
		return os.ReadFile(file)
	}

	r, err := statikFS.Open(GetStatikPath(file))
	if err != nil {
		return nil, err
	}
	return io.ReadAll(r)
}

func GetFileDataMust(file string) []byte {
	fileData, err := GetFileData(file)
	if err != nil {
		tool.PanicErr(err)
	}
	return fileData
}

func SetGinStatic(r *gin.Engine, relativePath string, root string) {
	if config.Config.IsDevelopment {
		r.Static(relativePath, root)
		return
	}
	r.StaticFS(relativePath, OpenDir(root))

}

func SetGinStaticFile(r *gin.Engine, relativePath string, filePath string) {
	if config.Config.IsDevelopment {
		r.StaticFile(relativePath, filePath)
		return
	}
	statikPath := GetStatikPath(filePath)
	r.GET(relativePath, func(ctx *gin.Context) {
		ctx.FileFromFS(statikPath, GetFileSystem())
	})
}

// func GetFile() {
// 	r, err := statikFS.Open("/index.html")
// 	if err != nil {
// 		tool.PanicErr(err)
// 	}
// 	defer r.Close()
// 	contents, err := ioutil.ReadAll(r)
// 	if err != nil {
// 		tool.PanicErr(err)
// 	}

// 	fmt.Println(string(contents))
// }
