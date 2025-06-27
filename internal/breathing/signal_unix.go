//go:build unix

package breathing

import (
	"os"
	"syscall"
)

var sigint = []os.Signal{syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT}

func sigexit(s os.Signal) {
	os.Exit(128 + int(s.(syscall.Signal)))
}
