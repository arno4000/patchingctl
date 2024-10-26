package patching

import (
	"bufio"
	"io"
	"os/exec"
	"time"

	"github.com/sirupsen/logrus"
)

func ExecuteCommand(command string) error {
	cmd := exec.Command("sh", "-c", command)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	// Function to log output from a pipe
	logOutput := func(pipe io.ReadCloser, logFunc func(args ...interface{})) {
		scanner := bufio.NewScanner(pipe)
		for scanner.Scan() {
			logFunc(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			logrus.Error("Error reading from pipe: ", err)
		}
	}

	// Log stdout and stderr in separate goroutines
	go logOutput(stdout, logrus.Info)
	go logOutput(stderr, logrus.Error)

	time.Sleep(10 * time.Millisecond)

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}
