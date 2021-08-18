package worker

import (
	"Mercurius/common"
	"fmt"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
	"unicode/utf8"
)

type WebSocketMaster struct {
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:    4096,
	WriteBufferSize:   4096,
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (w WebSocketMaster) Run(i int) error {
	http.HandleFunc("/", msgHandle)
	log.Infof("启动成功,端口号%d",common.GetConfig().Server.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", common.GetConfig().Server.Port), nil)


	return err
}

func msgHandle(w http.ResponseWriter, r *http.Request) {
	var writeMessage, writePrepared bool
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade:", err)
		return
	}
	defer conn.Close()
	for {
		mt, b, err := conn.ReadMessage()
		if err != nil {
			if err != io.EOF {
				log.Println("NextReader:", err)
			}
			return
		}
		if mt == websocket.TextMessage {
			if !utf8.Valid(b) {
				conn.WriteControl(websocket.CloseMessage,
					websocket.FormatCloseMessage(websocket.CloseInvalidFramePayloadData, ""),
					time.Time{})
				log.Println("ReadAll: invalid utf8")
			}
		}
		if writeMessage {
			if !writePrepared {
				err = conn.WriteMessage(mt, b)
				if err != nil {
					log.Println("WriteMessage:", err)
				}
			} else {
				pm, err := websocket.NewPreparedMessage(mt, b)
				if err != nil {
					log.Println("NewPreparedMessage:", err)
					return
				}
				err = conn.WritePreparedMessage(pm)
				if err != nil {
					log.Println("WritePreparedMessage:", err)
				}
			}
		} else {
			w, err := conn.NextWriter(mt)
			if err != nil {
				log.Println("NextWriter:", err)
				return
			}
			if _, err := w.Write(b); err != nil {
				log.Println("Writer:", err)
				return
			}
			if err := w.Close(); err != nil {
				log.Println("Close:", err)
				return
			}
		}
	}
}

func (w WebSocketMaster) SendData2Client(transmissionStruct common.TransmissionStruct) (int, error) {
	return 0, nil
}
