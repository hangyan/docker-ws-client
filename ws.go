package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"

	"log"
)

func UseXNet(url string, cmd string) {
	origin := "http://localhost"
	fmt.Printf("URL : %s\n", url)
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	message := []byte(cmd + "\n")
	fmt.Printf("Send: %s\n", message)

	if _, err := ws.Write(message); err != nil {
		log.Fatal(err)
	}

	for {
		var msg = make([]byte, 512)
		if _, err = ws.Read(msg); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", msg)
	}
}
