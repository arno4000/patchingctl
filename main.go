package main

import (
	"github.com/arno4000/patchingctl/pkg/patching"
	"github.com/sirupsen/logrus"
)

func main() {
	// utils.InitLogger()
	err := patching.UpdateOS()
	if err != nil {
		logrus.Errorln(err)
	}
	if patching.IsDockerInstalled() {
		err = patching.PatchDocker()
		if err != nil {
			logrus.Errorln(err)
		}
	}
	err = patching.Reboot()
	if err != nil {
		logrus.Errorln(err)
	}
}
