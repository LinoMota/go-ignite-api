package util

import (
	"os"
	"fmt"
	"github.com/americanas-go/config"
)

func loadEnv() {

	fmt.Println("Loading app enviroment...")

	var env = os.Getenv("GO_ENV")

	if env == "" {
		env = "dev"
	}

	fmt.Println("Using enviroment: ", env)

	var filename = fmt.Sprintf("./config/%s.yaml", env)

	os.Setenv("CONF", filename)

	config.Load()
}
