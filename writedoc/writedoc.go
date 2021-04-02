package writedoc

import (
	"os/user"
)

func Write(text string) {

}

func getHomeDir() string {
	username, _ := user.Current()
	return username.HomeDir
}
