package upload

import (
	"github.com/gorilla/websocket"
	"io"
)

type UploadRequest struct {
	filename string
	file     io.Reader
	conn     *websocket.Conn
}
