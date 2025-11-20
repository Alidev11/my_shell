package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var Path = os.Getenv("PATH")

var DIRS = strings.Split(Path, ":")

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, "$ ")
		scanner.Scan()
		cmd := scanner.Text()
		cmd = strings.TrimSpace(cmd)
		
		err := RunCmd(cmd)

		if err != nil {
			fmt.Println(err)
			continue
		}
	}

}
