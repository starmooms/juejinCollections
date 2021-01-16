package middleware

import (
	"fmt"
	"juejinCollections/logger"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"

	// "github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
	// "github.com/sirupsen/logrus"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 连接是否断开
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				reset := "\x1b[0m"

				if brokenPipe {
					logger.Logger.Errorf("[GIN Panic Recover]:\n%s\n%s%s", string(httpRequest), err, reset)
				} else {
					// headers := strings.Split(string(httpRequest), "\r\n")
					// for idx, header := range headers {
					// 	current := strings.Split(header, ":")
					// 	if current[0] == "Authorization" {
					// 		headers[idx] = current[0] + ": *"
					// 	}
					// }
					logger.Logger.Errorf("[GIN Panic Recover]:\n%s\n%s", string(httpRequest), err)
				}

				if brokenPipe {
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
				} else {
					c.AbortWithStatus(http.StatusInternalServerError)
				}

			}
		}()
		c.Next()
	}
}
