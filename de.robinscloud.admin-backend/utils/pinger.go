package utils

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func PingHost(host string) (bool, error) {
	out, _ := exec.Command("ping", host).Output()
	if strings.Contains(string(out), "nicht finden") {
		fmt.Println("error ", string(out))
		return false, errors.New("host in unreachable")
	} else {
		fmt.Println("success", string(out))
		return true, nil
	}
}
