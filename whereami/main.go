package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/shawnstephens517/whereami-go/detector"
)

func main() {
	// Get the current working directory.
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Detect the operating system.
	osName := runtime.GOOS
	fmt.Printf("Running on: %s\n", osName)

	detect(osName)

	// Print the current working directory and operating system.
	fmt.Printf("Current working directory: %s\n", dir)
}

// Detect identifies the operating system and runs the appropriate detection function.
func detect(os string) {
	switch os {
	// Detect the operating system.
	case "windows":
		fmt.Println("Running WhereAMI on Windows")
		detector.WindowsDetect()
	case "linux":
		fmt.Println("Running WhereAMI on Linux")
		detector.LinuxDetect()
	case "darwin":
		fmt.Println("Running WhereAMI on MacOS")
	default:
		fmt.Println("Unknown operating system")
	}

}
