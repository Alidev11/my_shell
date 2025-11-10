package main

import (
	"fmt"
	"os"
)


func main() {
	// TODO: Uncomment the code below to pass the first stage
	var cmd string
	for {
		fmt.Fprint(os.Stdout, "$ ")
		fmt.Scanln(&cmd)
		fmt.Fprint(os.Stdout, cmd + ": command not found\n")
	}
	

}
