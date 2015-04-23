package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/net/websocket"
)

const (
	WS_BASE_URL = "ws://localhost:4243"
	ORIGIN      = "http://localhost"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage : <program>  <contain id> <cmd>")
		return
	}

	id := os.Args[1]
	cmd := os.Args[2]

	wsURL := fmt.Sprintf("%s/v1.17/containers/%s/attach/ws?logs=0&stream=1&stdin=1&stdout=1&stderr=1", WS_BASE_URL, id)
	useXNet(wsURL, cmd)
}

func useXNet(url string, cmd string) {

	fmt.Printf("URL : %s\n", url)
	ws, err := websocket.Dial(url, "", ORIGIN)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	message := []byte(cmd + "\n")
	fmt.Printf("Send: %s\n", message)

	if _, err := ws.Write(message); err != nil {
		log.Fatal(err)
	}

	ws.SetReadDeadline(time.Now().Add(1 * time.Second))

	for {
		var msg = make([]byte, 512)
		var n int
		if n, err = ws.Read(msg); err != nil {
			break // usually is time out
		}
		fmt.Printf("%s", msg[:n])
		msg = nil
	}
}
