package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetPackages() string {
	d := GetDistro()

	switch {
	case strings.Contains(d, "arch"), strings.Contains(d, "cachy"), strings.Contains(d, "artix"), strings.Contains(d, "artix"):
		return runner("pacman -Qq", "pacman")

	case strings.Contains(d, "void"):
		return runner("xbps-query", "xbps")

	case strings.Contains(d, "alpine") || strings.Contains(d, "chimera"):
		return runner("apk list -I", "apk")

	case strings.Contains(d, "ubuntu"):
		return runner("apt list --installed", "apt")

	case strings.Contains(d, "kiss"):
		return runner("kiss list", "kiss")

	default:
		return "sorry"
	}
}

func runner(commands string, manager string) string {
	out, _ := exec.Command("sh", "-c", commands+" | wc -l").Output()

	return fmt.Sprintf("󰏓 %s packages (%s)", strings.TrimSpace(string(out)), manager)
}
