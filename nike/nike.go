package nike

import (
	"github.com/dev-appmonsters/dicemix-light-client/utils"
)

// NIKE - The main interface for Non-interactive Key Exchange (NIKE).
type NIKE interface {
	GenerateKeys(*utils.State, int)
	DeriveSharedKeys(*utils.State)
}
