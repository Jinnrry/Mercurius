package common

import "fmt"

const VERSION = "Version 0.1.1"

func PrintVersion(source string) {
	fmt.Printf("%s %s\n", source, VERSION)
}
