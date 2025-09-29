package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetPackages() string {
	managers := map[string]string{
		"pacman":  "pacman -Qq",
		"xbps":    "xbps-query -l",
		"apk":     "apk list -I",
		"apt":     "apt list --installed",
		"kiss":    "kiss list",
		"rpm":     "rpm -qa",
		"flatpak": "flatpak list",
	}

	results := []string{}

	for mgr, cmd := range managers {
		if isExecutable(mgr) {
			count := getPackageCount(cmd)
			results = append(results, fmt.Sprintf("󰏓 %s packages (%s)", count, mgr))
		}
	}

	if len(results) == 0 {
		return "sorry"
	}

	return strings.Join(results, " | ")
}

func getPackageCount(cmd string) string {
	out, err := exec.Command("sh", "-c", cmd+" | wc -l").Output()
	if err != nil {
		return "0"
	}
	return strings.TrimSpace(string(out))
}

func isExecutable(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
