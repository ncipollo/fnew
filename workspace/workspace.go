package workspace

import (
	"os/user"
	"path"
	"fmt"
)

func Directory() string {
	currentUser, err := user.Current()
	if err != nil {
		panic(fmt.Sprintf("Unable to get current user! Error: %s", err))
	}
	return path.Join(currentUser.HomeDir, ".fnew")
}
