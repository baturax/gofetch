package main

import (
	"fmt"
	"log"
	"os/exec"
	"os/user"
	"strings"

	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

func main() {
	iconLines := strings.Split(getDistro(), "\n")

	// Burada sistem bilgilerini satırlara bölüyoruz
	infoLines := []string{
		strings.TrimSpace(getUser()),
		strings.TrimSpace(getSystem()),
		strings.TrimSpace(getKernel()),
		strings.TrimSpace(getMem()),
		strings.TrimSpace(getUptime()),
		strings.TrimSpace(getPackages()),
	}

	// Kaç satır olduğuna bakalım (ikon veya bilgi satırı eksikse boş satır ekle)
	maxLines := len(iconLines)
	if len(infoLines) > maxLines {
		maxLines = len(infoLines)
	}

	// Eksik satır varsa boş satır ekleyelim
	for len(iconLines) < maxLines {
		iconLines = append(iconLines, "")
	}
	for len(infoLines) < maxLines {
		infoLines = append(infoLines, "")
	}

	// Satır satır yan yana birleştirip yazdırıyoruz
	for i := 0; i < maxLines; i++ {
		fmt.Printf("%-20s  %s\n", iconLines[i], infoLines[i])
	}
}

func getPackages() string {
	h, _ := host.Info()
	if h.Platform == "arch" {

		out, _ := exec.Command("pacman", "-Qq").Output()
		o := string(out[:])
		a := len(strings.Fields(o))
		return fmt.Sprintln("󰏖 ",a,   "pckages (pacman)")

	} else if h.Platform == "alpine" {

		out, _ := exec.Command("apk", "list", "-I").Output()
		o := string(out[:])
		a := len(strings.Split(o, "\n"))
		return fmt.Sprintln("󰏖 ",a,   "ackages (apk)")
	}
	return ""
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

	return fmt.Sprintf("  %v days, %v hours, %v minutes\n", day, hour, minute)
}

func getMem() string {
	m, err := mem.VirtualMemory()

	if err != nil {
		fmt.Println(err.Error())
	}

	usedMem := m.Used / 1048576
	totalMem := m.Total / 1048576

	return fmt.Sprintf("  %v / %v\n", usedMem, totalMem)
}

func getKernel() string {
	h, err := host.Info()

	if err != nil {
		fmt.Println(err.Error())
	}

	platform := h.Platform
	kernel := h.KernelVersion

	return fmt.Sprintf("  %v %v\n", platform, kernel)
}

func getSystem() string {
	h, err := host.Info()

	if err != nil {
		fmt.Println(err.Error())
	}

	platform := h.Platform
	os := h.OS
	arch := h.KernelArch

	return fmt.Sprintf("  %v %v %v\n", platform, os, arch)
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
	return fmt.Sprintf("%v@%v\n", user, hostName)
}

func getDistro() string {
	h, _ := host.Info()

	distro := h.Platform

	if strings.Contains(distro, "arch") {
		return `
      /\\          
     /  \\         
    /    \\        
   /      \\       
  /   ,,   \\      
 /   |  |   \\     
/_-''    ''-_\\    
`
	} else if strings.Contains(distro, "void") {
		return `
    _______      
 _ \______ -     
| \  ___  \ |    
| | /   \ | |    
| | \___/ | |    
| \______ \_|    
 -_______\       
`
	} else if strings.Contains(distro, "alpine") {
		return `
      /\\
     /  \\
    / /\\ \\  /\\
   / /  \\ \\/  \\
  / /    \\ \\/\\ \\
 /_/ /_|    \\_\\ \\_
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
