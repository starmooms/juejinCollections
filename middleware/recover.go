package middleware

import (
	"fmt"
	"juejinCollections/logger"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"reflect"
	"regexp"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	log := logger.GetLog()
	return func(c *gin.Context) {
		defer func() {
			if rErr := recover(); rErr != nil {
				// 连接是否断开
				var brokenPipe bool
				if ne, ok := rErr.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				// 生成堆载
				reflect.ValueOf(rErr).Type()
				var err error
				switch v := rErr.(type) {
				case error:
					rt := reflect.TypeOf(v)
					kind := rt.Kind()
					hasStack := false
					if kind == reflect.Ptr {
						rt = rt.Elem()
						kind = rt.Kind()
					}

					if kind == reflect.Struct {
						_, hasStack = rt.FieldByName("stack")
					}
					if hasStack {
						err = v
					} else {
						err = errors.NewWithDepth(2, v.Error())
					}
				default:
					err = errors.New(fmt.Sprintf("%v", rErr))
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, true)
				httpStr := string(httpRequest)
				reg, regErr := regexp.Compile("(\\r\\n)+$")
				if regErr == nil {
					httpStr = reg.ReplaceAllString(httpStr, "")
				}
				if brokenPipe {
					log.Errorf("[GIN Panic Recover]:\n%s\n%s%+v", httpStr, err)
				} else {
					// headers := strings.Split(string(httpRequest), "\r\n")
					// for idx, header := range headers {
					// 	current := strings.Split(header, ":")
					// 	if current[0] == "Authorization" {
					// 		headers[idx] = current[0] + ": *"
					// 	}
					// }
					log.Errorf("[GIN Panic Recover]:\n%s\n%+v", httpStr, err)
				}

				if brokenPipe {
					c.Error(err.(error))
					c.Abort()
				} else {
					c.AbortWithStatus(http.StatusInternalServerError)
				}
			}
		}()
		c.Next()
	}
}
