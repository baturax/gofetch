package mods

import (
	"fmt"

	"os/exec"
	"strings"
)

func GetPackages() string {
	if strings.Contains(getDistro(), "arch") || strings.Contains(getDistro(), "cachy") {
		return packageRunner("pacman -Qq", 1, "pacman")

	} else if strings.Contains(getDistro(), "void") {
		return packageRunner("xbps-query", 1, "xbps")

	} else if strings.Contains(getDistro(), "alpine") {
		return packageRunner("apk list -I", 1, "apk")

	} else if strings.Contains(getDistro(), "ubuntu") {
		return packageRunner("apt list --installed", 2, "apt")

	} else {
		return fmt.Sprintf("sorry")

	}
}

func packageRunner(commands string, mins int, manager string) string {

	c, _ := exec.Command("sh", "-c", commands).Output()
	o := string(c[:])
	b := strings.Split(o, "\n")
	return fmt.Sprintln("󰏓 ", len(b)-mins, "packages", manager)
}
