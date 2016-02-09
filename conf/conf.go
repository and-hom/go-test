package conf

import (
	"os"
	"fmt"
)

func init() {
	homedir := os.Getenv("HOME")
	fmt.Printf("HOME=%s\n", homedir)
}

func SayHello() {
	fmt.Printf("Hello!\n")
}
