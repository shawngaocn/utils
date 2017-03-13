package runner

import (
	"os/exec"
	"strings"
)

func RunGoFile(path string, args string) (result []byte, err error) {
	if args != "" {
		args = "run " + path + " " + args
	} else {
		args = "run " + path
	}
	argArray := strings.Split(args, " ")
	cmd := exec.Command("go", argArray...)
	return cmd.Output()
}
