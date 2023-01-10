package websocket

import (
	// "flag"
	// "fmt"
	"juejinCollections/logger"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var log = logger.Logger
var upgrader = websocket.Upgrader{} // use default option

func echo(ctx *gin.Context) {
	w, r := ctx.Writer, ctx.Request
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorf("upgrade:", err)
		return
	}

	InitCollectWs(conn)

	// defer conn.Close()

	// for {
	// 	mt, message, err := conn.ReadMessage()
	// 	if err != nil {
	// 		// 如果是关闭错误则正常关闭，不打印错误
	// 		// 本来应该前端发送close帧，后端停止ReadMessage
	// 		// 如果前端直接close，没有状态码是 1005 websocket.CloseNoStatusReceived，正常关闭状态码 1000 websocket.CloseNormalClosure
	// 		// https://github.com/gorilla/websocket/issues/575
	// 		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure, websocket.CloseNormalClosure, websocket.CloseNoStatusReceived) {
	// 			log.Errorf("error: %v", err)
	// 		} else {
	// 			if e, ok := err.(*websocket.CloseError); ok {
	// 				log.Debug("webstock close code ", e.Code)
	// 			} else {
	// 				log.Debug("webstock close no websocket.CloseError")
	// 			}
	// 		}
	// 		break
	// 	}

	// 	if string(message) == "1" {
	// 		conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "server close"))
	// 		conn.Close()
	// 		break
	// 	}

	// 	log.Printf("recv:%s", message)
	// 	err = conn.WriteMessage(mt, message)
	// 	if err != nil {
	// 		log.Println("write:", err)
	// 		break
	// 	}
	// }
}

// func home(c *gin.Context) {
// 	if r.URL.Path != "/" {
// 		http.Error(w, "Not found", http.StatusNotFound)
// 		return
// 	}
// 	if r.Method != http.MethodGet {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}
// }

func Start(r *gin.Engine) {
	// var addr = flag.String("addr", fmt.Sprintf(":%d", port), "http service address")
	// flag.Parse()
	r.GET("/echo", echo)
	// r.Run(*addr)
}
