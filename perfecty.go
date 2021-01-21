package perfecty

import (
	"github.com/rwngallego/perfecty-push/internal"
)

const filePath = "configs/internal.yml"

// Start Setup and start the push server
func Start() (err error) {
	if err = internal.LoadConfig(filePath); err != nil {
		return
	}

	if err = internal.LoadLogging(); err != nil {
		return
	}

	if err = internal.StartServer(); err != nil {
		return
	}

	return
}
