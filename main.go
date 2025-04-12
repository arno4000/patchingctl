package main

import (
	"github.com/arno4000/patchingctl/pkg/patching"
	"github.com/arno4000/patchingctl/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	err := utils.LoadConfig("/etc/patching.yaml")
	if err != nil {
		logrus.Fatalln(err)
	}
	utils.InitLogger()
	err = patching.UpdateOS()
	if err != nil {
		logrus.Fatalln(err)
	}
	if patching.IsDockerInstalled() {
		err = patching.PatchDocker()
		if err != nil {
			logrus.Fatalln(err)
		}
	}
	if viper.GetBool("reboot") {
		err = patching.Reboot()
		if err != nil {
			logrus.Fatalln(err)
		}
	}
}
