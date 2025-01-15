package detector

import (
	"fmt"
	"os"
	"runtime"
)

// MacOSDetect detects if the operating system is macOS and prints the hostname.
func MacOSDetect() {

	if runtime.GOOS == "darwin" {
		fmt.Println("macOS")
	} else {
		fmt.Println("Not macOS")
	}

	println(os.Hostname())

}
