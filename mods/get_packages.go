package mods

import (
	"fmt"

	"os/exec"
	"strings"
)

func GetPackages() string {
	if strings.Contains(getDistro(), "arch") {

		c, _ := exec.Command("pacman", "-Qq").Output()
		o := string(c[:])
		b := strings.Split(o, "\n")
		return fmt.Sprintln("󰏓", len(b)-1, "packages (pacman)")

	} else if strings.Contains(getDistro(), "void") {

		c, _ := exec.Command("xbps-query", "-l").Output()
		o := string(c[:])
		b := strings.Split(o, "\n")
		return fmt.Sprintln("󰏓", len(b)-1, "packages (xbps)")

	} else if strings.Contains(getDistro(), "alpine") {

		c, _ := exec.Command("apk", "list", "-I").Output()
		o := string(c[:])
		b := strings.Split(o, "\n")
		return fmt.Sprintln("󰏓", len(b)-1, "packages (apk)")

	} else if strings.Contains(getDistro(), "ubuntu") {

		c, _ := exec.Command("apt", "list", "--installed").Output()
		o := string(c[:])
		b := strings.Split(o, "\n")
		return fmt.Sprintln("󰏓", len(b)-2, "packages (apt)")

	} else {
		return fmt.Sprintf("sorry")
	}
}
