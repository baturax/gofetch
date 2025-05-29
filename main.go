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
	fons := []string{
		getUser(),
		getSystem(),
		getKernel(),
		getMem(),
		getUptime(),
		getShell(),
		getDesktop(),
		getPackages(),
	}

	mL := len(getIcon())
	if l := len(fons); l > mL {
		mL = l
	}

	for i := range mL {
		icon := ""
		info := ""

		if i < len(getIcon()) {
			icon = getIcon()[i]
		} else {
			icon = "                  "
		}

		if i < len(fons) {
			info = fons[i]
		}

		fmt.Printf("%s  %s\n", icon, info)
	}
}

func getPackages() string {
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

func getIcon() []string {
	distro := strings.ToLower(getDistro())
	switch {
	case strings.Contains(distro, "arch"):
		return []string{
			"      /\\          ",
			"     /  \\         ",
			"    /    \\        ",
			"   /      \\       ",
			"  /   ,,   \\      ",
			" /   |  |   \\     ",
			"/_-''    ''-_\\    ",
		}
	case strings.Contains(distro, "void"):
		return []string{
			"    _______        ",
			" _ \\______ -      ",
			"| \\  ___  \\ |     ",
			"| | /   \\ | |     ",
			"| | \\___/ | |     ",
			"| \\______ \\_|     ",
			" -_______\\        ",
		}
	case strings.Contains(distro, "alpine"):
		return []string{
			"      /\\          ",
			"     /  \\         ",
			"    / /\\ \\  /\\    ",
			"   / /  \\ \\/  \\   ",
			"  / /    \\ \\ /\\   ",
			" / / /|   \\ \\ \\   ",
			"/_/ /_|    \\_\\_\\  ",
		}
	case strings.Contains(distro, "ubuntu"):
		return []string{
			"         _        ",
			"     ---(_)       ",
			" _/  ---  \\       ",
			"(_) |   |         ",
			"  \\  --- _/       ",
			"     ---(_)       ",
		}
	default:
		return []string{
			"    ___            ",
			"   (.. \\          ",
			"   (<> |           ",
			"  //  \\ \\        ",
			" ( |  | /|         ",
			"_/\\ __)/_)        ",
			"\\/-____\\/        ",
		}
	}
}
