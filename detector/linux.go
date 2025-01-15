package detector

import (
	"fmt"
	"os"
	"runtime"
)

// LinuxDetect prints the operating system and hostname if running on Linux.
func LinuxDetect() {

	if runtime.GOOS == "linux" {
		fmt.Println("Linux")
	} else {
		fmt.Println("Not Linux")
	}
	println(os.Hostname())

}
