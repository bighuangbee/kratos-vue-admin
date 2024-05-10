package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

func main() {
	wsUrl := "ws://zj-cn-live-comet.chat.bilibili.com:2244/sub"

	// 连接 WebSocket
	conn, resp, err := websocket.DefaultDialer.Dial(wsUrl, nil)
	if err != nil {
		fmt.Printf("dial error: %v\n", err)
		return
	}

	defer func(conn *websocket.Conn) {
		// 关闭底层 TCP 网络连接
		err := conn.Close()
		if err != nil {
			fmt.Printf("close tcp conn error: %v\n", err)
		}
	}(conn)

	// 输出响应信息
	fmt.Printf("ws resp: Status=%s, StatusCode=%d, Proto=%s\n",
		resp.Status, resp.StatusCode, resp.Proto)

	err1 := conn.WriteJSON(map[string]interface{}{
		"roomid": 26896215,
		"key":    "xZHL34ZShTlBCb1glXmKwz5_p7iE8CQ3ZP8hN4DNkUhAJ8p1i8yYCaVEfkmDXS9yE9yd5WnjShNDDRt6cvE_e5btD0Rb5LYE36krV3DjZm1mUpS4BSLLKRFgM3dfdRzgvxA2PZzP8oCxlNdBZuGqGLbmVHaJGeXte35Djn2MA_pSF1wNUAJ0p2q2NA==",
	})
	fmt.Println("发送认证包 ", err1)

	for {
		// 读取 (Text/Binary) 数据消息 (接收到 Close 消息或连接异常断开时, 此方法结束阻塞并返回错误)。
		// msgType 的值为 websocket.TextMessage 或 websocket.BinaryMessage,
		// msg 是 []byte 类型, 如果死 Text 消息, 则为 UTF-8 编码的文本。
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("read error: %v\n", err)
			//break
		}

		// 打印读取到的数据消息
		fmt.Printf("read server msg: msgType=%d, msg=%s\n", msgType, string(msg))

		time.Sleep(time.Second)
	}

}
