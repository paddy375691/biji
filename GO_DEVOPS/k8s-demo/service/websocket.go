package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"k8s.io/client-go/tools/remotecommand"
)

const END_OF_TRANSMISSION = "\u0004"

// TerminalMessage is the messaging protocol between ShellController and TerminalSession.
type TerminalMessage struct {
	Operation string `json:"operation"`
	Data      string `json:"data"`
	Rows      uint16 `json:"rows"`
	Cols      uint16 `json:"cols"`
}

var upgrader = func() websocket.Upgrader {
	upgrader := websocket.Upgrader{}
	upgrader.HandshakeTimeout = time.Second * 2
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	return upgrader
}()

// TerminalSession implements PtyHandler
type TerminalSession struct {
	wsConn   *websocket.Conn
	sizeChan chan remotecommand.TerminalSize
	doneChan chan struct{}
	tty      bool
}

// NewTerminalSessionWs create TerminalSession
func NewTerminalSessionWs(conn *websocket.Conn) *TerminalSession {
	return &TerminalSession{
		wsConn:   conn,
		tty:      true,
		sizeChan: make(chan remotecommand.TerminalSize),
		doneChan: make(chan struct{}),
	}
}

// NewTerminalSession create TerminalSession
func NewTerminalSession(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*TerminalSession, error) {
	conn, err := upgrader.Upgrade(w, r, responseHeader)
	if err != nil {
		return nil, err
	}
	session := &TerminalSession{
		wsConn:   conn,
		tty:      true,
		sizeChan: make(chan remotecommand.TerminalSize),
		doneChan: make(chan struct{}),
	}
	return session, nil
}

// Done done, must call Done() before connection close, or Next() would not exits.
func (t *TerminalSession) Done() {
	close(t.doneChan)
}

// Next called in a loop from remotecommand as long as the process is running
func (t *TerminalSession) Next() *remotecommand.TerminalSize {
	select {
	case size := <-t.sizeChan:
		return &size
	case <-t.doneChan:
		return nil
	}
}

// Stdin ...
func (t *TerminalSession) Stdin() io.Reader {
	return t
}

// Stdout ...
func (t *TerminalSession) Stdout() io.Writer {
	return t
}

// Stderr ...
func (t *TerminalSession) Stderr() io.Writer {
	return t
}

// Read called in a loop from remotecommand as long as the process is running
func (t *TerminalSession) Read(p []byte) (int, error) {
	_, message, err := t.wsConn.ReadMessage()

	if err != nil {
		log.Printf("read message err: %v", err)
		return copy(p, END_OF_TRANSMISSION), err
	}
	var msg TerminalMessage
	if err := json.Unmarshal([]byte(message), &msg); err != nil {
		log.Printf("read parse message err: %v", err)
		// return 0, nil
		return copy(p, END_OF_TRANSMISSION), err
	}

	switch msg.Operation {
	case "stdin":
		return copy(p, msg.Data), nil
	case "resize":
		t.sizeChan <- remotecommand.TerminalSize{Width: msg.Cols, Height: msg.Rows}
		return 0, nil
	case "ping":
		return 0, nil
	default:
		log.Printf("unknown message type '%s'", msg.Operation)
		// return 0, nil
		return copy(p, END_OF_TRANSMISSION), fmt.Errorf("unknown message type '%s'", msg.Operation)
	}
}

// Write called from remotecommand whenever there is any output
func (t *TerminalSession) Write(p []byte) (int, error) {
	msg, err := json.Marshal(TerminalMessage{
		Operation: "stdout",
		Data:      string(p),
	})
	if err != nil {
		log.Printf("write parse message err: %v", err)
		return 0, err
	}
	if err := t.wsConn.WriteMessage(websocket.TextMessage, msg); err != nil {
		log.Printf("write message err: %v", err)
		return 0, err
	}
	return len(p), nil
}

// Tty ...
func (t *TerminalSession) Tty() bool {
	return t.tty
}

// Close close session
func (t *TerminalSession) Close() error {
	return t.wsConn.Close()
}