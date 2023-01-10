package websocket

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"juejinCollections/collectReq"
)

type wsData struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type collectWs struct {
	conn   *websocket.Conn
	send   chan interface{}
	isConn bool
	logId  int
}

func (c *collectWs) close() {
	if c.isConn {
		collectReq.DelRunLog(c.logId)
		c.conn.Close()
		close(c.send)
		c.isConn = false
	}
}

func (c *collectWs) read() {
	defer c.close()

	for {
		data := &wsData{}
		err := c.conn.ReadJSON(data)
		if err != nil {
			// 如果是关闭错误则正常关闭，不打印错误
			// 本来应该前端发送close帧，后端停止ReadMessage
			// 如果前端直接close，没有状态码是 1005 websocket.CloseNoStatusReceived，正常关闭状态码 1000 websocket.CloseNormalClosure
			// https://github.com/gorilla/websocket/issues/575
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure, websocket.CloseNormalClosure, websocket.CloseNoStatusReceived) {
				log.Errorf("collectWs error: %v", err)
			} else {
				if e, ok := err.(*websocket.CloseError); ok {
					log.Debug("webstock close code ", e.Code)
				} else {
					log.Debug("webstock close no websocket.CloseError")
				}
			}
			break
		}

		if data.Type == "action" {
			if data.Data == "run" {
				c.openAction()
			}
		}
	}
}

func (c *collectWs) write() {
	defer c.close()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				return
			}
			c.conn.WriteJSON(message)
		}
	}
}

func (c *collectWs) openAction() {
	if collectReq.HasRunAction {
		c.send <- &gin.H{
			"type": "tip",
			"msg":  "当前正在同步收藏集，不能重复触发",
		}
	}
	c.WriteLog("开始同步收藏集...")
	go collectReq.Run()
}

func (c *collectWs) WriteLog(data string) {
	if c.isConn {
		c.send <- &gin.H{
			"type": "log",
			"data": data,
		}
	}
}

func InitCollectWs(coon *websocket.Conn) {
	c := &collectWs{
		send:   make(chan interface{}, 10),
		conn:   coon,
		isConn: true,
	}
	c.logId = collectReq.SetRunLog(c.WriteLog)
	go c.read()
	go c.write()

	// go func() {
	// 	for {
	// 		time.Sleep(1 * time.Second)
	// 		c.send <- &gin.H{
	// 			"a": 1,
	// 		}
	// 	}
	// }()
}
