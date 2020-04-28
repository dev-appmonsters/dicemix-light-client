package server

import (
	"flag"
	"net/url"
	"time"

	"github.com/dev-appmonsters/dicemix-light-client/dc"
	"github.com/dev-appmonsters/dicemix-light-client/messages"
	"github.com/dev-appmonsters/dicemix-light-client/nike"
	"github.com/dev-appmonsters/dicemix-light-client/utils"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"

	log "github.com/sirupsen/logrus"
)

// server configurations
var addr = flag.String("addr", "localhost:8082", "http service address")
var dialer = websocket.Dialer{} // use default options

// Exposed interfaces
var iNike nike.NIKE
var iDcNet dc.DC

type connection struct {
	Server
}

// NewConnection creates a new Server instance
func NewConnection() Server {
	initialize()

	return &connection{}
}

// Register - requests to C_JOIN_REQUEST
func (c *connection) Register(state *utils.State) {
	var connection = connect()
	listener(connection, state)

	defer connection.Close()
}

// performs some basic initializations
func initialize() {
	// initailze exposed interfaes for further use
	iNike = nike.NewNike()
	iDcNet = dc.NewDCNetwork()
}

// connects to server and extablishes a web socket connection
func connect() *websocket.Conn {
	url := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	log.Info("Connecting to ", url.String())
	conn, _, err := dialer.Dial(url.String(), nil)
	checkError(err)
	log.Info("Connected to ", url.String())

	// Read the message from server with deadline of ResponseWait(30) seconds
	conn.SetReadDeadline(time.Now().Add(utils.ResponseWait * time.Second))

	return conn
}

// listens for responses from server side
func listener(c *websocket.Conn, state *utils.State) {
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Fatalf("Connection closed - %v", err)
		}

		response := &messages.GenericResponse{}
		err = proto.Unmarshal(message, response)
		checkError(err)

		// handles response and take further actions
		// based on response.Code
		handleMessage(c, message, response.Header.Code, state)
	}
}
