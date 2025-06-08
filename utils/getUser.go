package utils

import (
	"fmt"
	"os/user"
)

func GetUser() string {
	userr, err := user.Current()
	showError(err)

	user := userr.Username

	return fmt.Sprintf("%v@%v", user, getHost().Hostname)

}
