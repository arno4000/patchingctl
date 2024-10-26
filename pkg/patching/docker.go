package patching

import (
	"fmt"
	"os"
)

func PatchDocker() error {
	composeBaseDir := "/srv/compose"
	composeDirs, err := os.ReadDir(composeBaseDir)
	for _, composeDir := range composeDirs {
		if composeDir.IsDir() {
			composeFile := fmt.Sprintf("%s/%s/compose.yml", composeBaseDir, composeDir.Name())
			err := ExecuteCommand(fmt.Sprintf("docker compose --ansi never -f %s up -d --pull always 2>&1", composeFile))
			if err != nil {
				return err
			}
			err = ExecuteCommand("docker system prune -af")
			if err != nil {
				return err
			}
		}
	}
	if err != nil {
		return err
	}
	return nil
}

func IsDockerInstalled() bool {
	_, err := os.Stat("/usr/bin/docker")
	return !os.IsNotExist(err)
}
