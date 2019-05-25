package utils

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/mritd/caddybuilder/conf"

	"github.com/mitchellh/go-homedir"

	"github.com/sirupsen/logrus"
)

func CheckAndExit(err error) {
	if err != nil {
		if conf.Debug {
			logrus.Panic(err)
		} else {
			logrus.Fatal(err)
		}
	}
}

func ShortenString(str string, n int) string {
	if len(str) <= n {
		return str
	} else {
		return str[:n]
	}
}

func CheckCommand(cmd ...string) bool {
	if len(cmd) == 0 {
		logrus.Fatal("command is empty!")
	}
	return exec.Command(cmd[0], cmd[1:]...).Run() == nil
}

func GetGoPath() string {
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		home, err := homedir.Dir()
		CheckAndExit(err)
		goPath = filepath.Join(home, "gopath")
		logrus.Warnf("GOPATH not set, will auto create to [%s]", goPath)
		_, err = os.Stat(goPath)
		if err != nil {
			if os.IsNotExist(err) {
				err = os.MkdirAll(goPath, 0755)
				CheckAndExit(err)
			} else {
				CheckAndExit(err)
			}
		}

	}
	return goPath
}

func GetCaddyRepoPath() string {
	return filepath.Join(GetGoPath(), "src", "github.com", "mholt", "caddy")
}
