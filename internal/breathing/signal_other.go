//go:build !unix

package breathing

import (
	"os"
)

var sigint = []os.Signal{os.Interrupt, os.Kill}

func sigexit(s os.Signal) {
	os.Exit(1)
}
