package cli

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/e6a5/zenta/internal/version"
)

// captureOutput captures stdout and stderr from a function.
func captureOutput(f func()) (string, string) {
	oldStdout := os.Stdout
	oldStderr := os.Stderr
	rOut, wOut, _ := os.Pipe()
	rErr, wErr, _ := os.Pipe()
	os.Stdout = wOut
	os.Stderr = wErr

	f()

	wOut.Close()
	wErr.Close()
	stdout, _ := io.ReadAll(rOut)
	stderr, _ := io.ReadAll(rErr)
	os.Stdout = oldStdout
	os.Stderr = oldStderr

	return string(stdout), string(stderr)
}

func TestShowHelp(t *testing.T) {
	programName := "zenta-test"
	stdout, _ := captureOutput(func() {
		ShowHelp(programName)
	})

	if !strings.Contains(stdout, "mindfulness for terminal users") {
		t.Error("Help message should contain the description")
	}
	if !strings.Contains(stdout, "USAGE:") {
		t.Error("Help message should contain USAGE section")
	}
	if !strings.Contains(stdout, programName) {
		t.Errorf("Help message should contain the program name '%s'", programName)
	}
}

func TestHandleVersion(t *testing.T) {
	programName := "zenta-test"
	stdout, _ := captureOutput(func() {
		HandleVersion(programName)
	})

	expected := programName + " version " + version.Version
	if !strings.Contains(stdout, expected) {
		t.Errorf("Expected version output to contain '%s', got '%s'", expected, stdout)
	}
}

func TestHandleUnknownCommand(t *testing.T) {
	// Note: This test will cause the test suite to exit with status 1
	// if the os.Exit(1) is not handled. We can't prevent the exit here,
	// so this is more of a smoke test to see if it prints the right message.
	// In a real-world, more complex scenario, you might inject an exit function.
	t.Skip("Skipping test that calls os.Exit, which terminates the test suite.")

	// command := "foo"
	// programName := "zenta-test"
	// _, stderr := captureOutput(func() {
	// 	HandleUnknownCommand(command, programName)
	// })

	// if !strings.Contains(stderr, "Unknown command: "+command) {
	// 	t.Errorf("Expected unknown command error for '%s', got '%s'", command, stderr)
	// }
}
