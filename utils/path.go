package utils

import (
	"bytes"
	"log"
	"os"
	"os/exec"
)

func InteractiveToexec(commandName string, params []string) (string, bool) {
	cmd := exec.Command(commandName, params...)
	buf, err := cmd.Output()
	log.Println(cmd.Args)
	w := bytes.NewBuffer(nil)
	cmd.Stderr = w
	log.Printf("%s\n", w)
	if err != nil {
		log.Println("Error: <", err, "> when exec command read out buffer")
		return "", false
	} else {
		return string(buf), true
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
