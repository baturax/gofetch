package mods

import (
	"fmt"
	"log"
	"os/user"

	"github.com/shirou/gopsutil/v4/host"
)

func GetUser() string {
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
