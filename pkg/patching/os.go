package patching

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func DetectOS() (string, error) {
	if runtime.GOOS == "linux" {
		idLike, err := exec.Command("grep", "ID_LIKE", "/etc/os-release").Output()
		if err != nil {
			return "", err
		}
		idLikeStr := string(idLike)
		if strings.Contains(idLikeStr, "debian") {
			return "debian", nil
		} else if strings.Contains(idLikeStr, "rhel") {
			return "rhel", nil
		}
	} else {
		return "", fmt.Errorf("%s is unsupported", runtime.GOOS)
	}
	return "", fmt.Errorf("could not detect OS")
}

func Reboot() error {
	distro, err := DetectOS()
	if err != nil {
		return err
	}
	if distro == "debian" {
		if _, err := os.Stat("/var/run/reboot-required"); !os.IsNotExist(err) {
			return ExecuteCommand("reboot")
		}
	} else if distro == "rhel" {
		cmd := exec.Command("needs-restarting", "-r")
		err := cmd.Run()
		if err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				if exitError.ExitCode() == 1 {
					return ExecuteCommand("reboot")
				}
			}
		}
	}
	return nil
}
