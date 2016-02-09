package conf

import (
	"os"
	"fmt"
)

func init() {
	homedir := os.Getenv("USER_HOME")
	fmt.Printf("HOME=%s\n", homedir)
}

func SayHello() {
	fmt.Printf("Hello!\n")
}
