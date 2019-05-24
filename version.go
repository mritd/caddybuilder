package main

import (
	"encoding/base64"
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var bannerBase64 = "CiAgX19fICAgICAgICAgICAgICAgICAgICAgICAgICAgICwgX18gICAgICAgICAgICBfICAgICAgICAgICAgICAgICAgCiAvIChfKSAgICAgICAgIHwgICAgIHwgICAgICAgICAgL3wvICBcICAgICAgIG8gIHwgfCAgICB8ICAgICAgICAgICAgCnwgICAgICBfXywgICBfX3wgICBfX3wgICAgICAgICAgIHwgX18vICAgICAgICAgIHwgfCAgX198ICAgXyAgICxfICAgCnwgICAgIC8gIHwgIC8gIHwgIC8gIHwgIHwgICB8ICAgIHwgICBcfCAgIHwgIHwgIHwvICAvICB8ICB8LyAgLyAgfCAgCiBcX19fL1xfL3xfL1xfL3xfL1xfL3xfLyBcXy98LyAgIHwoX18vIFxfL3xfL3xfL3xfXy9cXy98Xy98X18vICAgfF8vCiAgICAgICAgICAgICAgICAgICAgICAgICAgIC98ICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgIFx8ICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCgo="

var versionTpl = `%s
Version: %s
Arch: %s
BuildDate: %s
CommitID: %s
`

var (
	Version   string
	BuildDate string
	CommitID  string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Long: `
Print version.`,
	Run: func(cmd *cobra.Command, args []string) {
		banner, _ := base64.StdEncoding.DecodeString(bannerBase64)
		fmt.Printf(versionTpl, banner, Version, runtime.GOOS+"/"+runtime.GOARCH, BuildDate, CommitID)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
