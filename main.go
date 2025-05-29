package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"

	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

func main() {
	fmt.Println(getIcon())
	fmt.Println(getUser())
	fmt.Println(getSystem())
	fmt.Println(getKernel())
	fmt.Println(getMem())
	fmt.Println(getUptime())
	fmt.Println(getShell())
	fmt.Println(getDesktop())
	getPackages()
}

func getPackages() {
	if strings.Contains(getDistro(), "arch") {

		c, _ := exec.Command("pacman", "-Qq").Output()
		o := string(c[:])
		b := strings.Split(o, "\n")
		fmt.Println(len(b) - 1)

	} else if strings.Contains(getDistro(), "void") {

		c, _ := exec.Command("xbps-query", "-l").Output()
		o := string(c[:])
		b := strings.Split(o, "\n")
		fmt.Println(len(b) - 1)

	} else if strings.Contains(getDistro(), "alpine") {

		c, _ := exec.Command("apk", "list", "-I").Output()
		o := string(c[:])
		b := strings.Split(o, "\n")
		fmt.Println(len(b) - 1)

	} else if strings.Contains(getDistro(), "ubuntu") {

		c, _ := exec.Command("apt", "list", "--installed").Output()
		o := string(c[:])
		b := strings.Split(o, "\n")
		fmt.Println(len(b) - 2)

	} else {

	}
}

func getDesktop() string {
	s := os.Getenv("XDG_CURRENT_DESKTOP")
	if s == "" {
		return fmt.Sprintf("  sorry %v", s)
	}
	return fmt.Sprintf("  %v", s)

}

func getShell() string {
	s := os.Getenv("SHELL")
	if s == "" {
		return fmt.Sprintf("󰲒 sorry %v", s)
	}

	return fmt.Sprintf("󰲒 %v", s)

}

func getUptime() string {
	h, err := host.Info()
	if err != nil {
		fmt.Println(err.Error())
	}

	uptime := h.Uptime

	day := uptime / 86400
	hour := (uptime % 86400) / 3600
	minute := (uptime % 3600) / 60

	return fmt.Sprintf("  %v days, %v hours, %v minutes", day, hour, minute)
}

func getMem() string {
	m, err := mem.VirtualMemory()

	if err != nil {
		fmt.Println(err.Error())
	}

	usedMem := m.Used / 1048576
	totalMem := m.Total / 1048576

	return fmt.Sprintf("  %v / %v", usedMem, totalMem)
}

func getKernel() string {
	h, err := host.Info()

	if err != nil {
		fmt.Println(err.Error())
	}

	platform := h.Platform
	kernel := h.KernelVersion

	return fmt.Sprintf("  %v %v", platform, kernel)
}

func getSystem() string {
	h, err := host.Info()

	if err != nil {
		fmt.Println(err.Error())
	}

	platform := h.Platform
	os := h.OS
	arch := h.KernelArch

	return fmt.Sprintf("  %v %v %v", platform, os, arch)
}

func getUser() string {
	u, uerr := user.Current()

	if uerr != nil {
		log.Fatal(uerr.Error())
	}

	h, herr := host.Info()

	if herr != nil {
		log.Fatal(herr.Error())
	}

	user := u.Username
	hostName := h.Hostname
	return fmt.Sprintf("%v@%v", user, hostName)
}

func getDistro() string {
	h, _ := host.Info()
	d := h.Platform
	return d
}

func getIcon() string {
	if strings.Contains(getDistro(), "arch") {
		return `
      /\
     /  \
    /    \
   /      \
  /   ,,   \
 /   |  |   \
/_-''    ''-_\
`
	} else if strings.Contains(getDistro(), "void") {
		return `
    _______
 _ \______ -
| \  ___  \ |
| | /   \ | |
| | \___/ | |
| \______ \_|
 -_______\
`
	} else if strings.Contains(getDistro(), "alpine") {
		return `
      /\
     /  \
    / /\ \  /\
   / /  \ \/  \
  / /    \ \/\ \
 / / /|   \ \ \ \
/_/ /_|    \_\ \_\
`
	} else if strings.Contains(getDistro(), "ubuntu") {
		return `
         _
     ---(_)
_/  ---  \
(_) |   |
  \  --- _/
     ---(_)
`
	}

	return `
    ___
   (.. \
   (<> |
  //  \ \
 ( |  | /|
_/\ __)/_)
\/-____\/
`
}
