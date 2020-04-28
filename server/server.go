package server

import (
	"github.com/dev-appmonsters/dicemix-light-client/utils"
)

// Server - The main interface to enable connection with server.
type Server interface {
	Register(*utils.State)
}
