package main

import (
	"fmt"
	"os"
)

func main() {

	id := os.Args[1]
	cmd := os.Args[2]
	wsURL := fmt.Sprintf("ws://localhost:4243/v1.17/containers/%s/attach/ws?logs=0&stream=1&stdin=1&stdout=1&stderr=1", id)
	UseXNet(wsURL, cmd)

}
