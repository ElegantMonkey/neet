package util

import (
	"github.com/ElegantMonkey/neet/config"
	"os/exec"
)

func LaunchPlayer(videoPath string) {
	cmd := exec.Command(config.Conf.Player, videoPath)
	cmd.Run()
}
