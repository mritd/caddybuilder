package utils

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/mitchellh/go-homedir"

	"github.com/sirupsen/logrus"
)

func CheckAndExit(err error) {
	if err != nil {
		logrus.Panic(err)
	}
}

func CheckGoCommand() bool {
	return exec.Command("go", "version").Run() == nil
}

func GetGoPath() string {
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		home, err := homedir.Dir()
		CheckAndExit(err)
		goPath = filepath.Join(home, "gopath")
		logrus.Warnf("GOPATH not exist, will auto create to [%s]", goPath)
		err = os.Mkdir(goPath, 0755)
		CheckAndExit(err)
	}
	return goPath
}

func GetCaddyRepoPath() string {
	return filepath.Join(GetGoPath(), "src", "github.com", "mholt", "caddy")
}
