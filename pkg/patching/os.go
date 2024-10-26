package patching

import (
	"fmt"
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
	return ExecuteCommand("reboot")
}
