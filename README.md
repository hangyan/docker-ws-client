# Docker Remote Attch Use Websocket 
This is a simple example of docker remote attach use websocket protocol.

## Install
`go get github.com/hangyan/docker-ws-client`

## Usage
![](https://raw.githubusercontent.com/hangyan/docker-ws-client/master/ws.png)

The first argument is the container id,and second is the command you want
to execute in that container

## Notes
The commands that has stream output may interfere with the later commmand's
output,like `top`,`ping`,so you need to restart the container after you execute
these commands.


## Others
You can also use a web browser to act as a websocket client too,there is
a chrome extension can do this : [Simple Websocket Client](https://chrome.google.com/webstore/detail/simple-websocket-client/pfdhoblngboilpfeibdedpjgfnlcodoo?hl=en)


![](https://raw.githubusercontent.com/hangyan/docker-ws-client/master/extension.png)


Notes:

1. You shoud add linebreak when you input the command
2. the example's (`ls -al`) output has some strange chars,they are terminal's
color control character.
3. Don't use `logs=1` or the output will be messed up.


