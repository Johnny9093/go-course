package cli

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

type (
	Config struct {
		Username string
		Password string
		Threads int
	}
)

func TestCli(t *testing.T) {
	config := &Config{}

	flag.StringVar(&config.Username, "username", "admin", "Some user")
	flag.StringVar(&config.Password, "password", "", "Your password")
	flag.IntVar(&config.Threads, "threads", 2, "# of threads")
	flag.Parse()

	fmt.Printf("num of args: %d\n", flag.NArg())
	fmt.Printf("num of flags: %d\n", flag.NFlag())

	// bad input, print help, exit with error
	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Printf("Config: %+v", config)
}
