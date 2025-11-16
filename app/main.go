package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var Path = os.Getenv("PATH")

var Arr = strings.Split(Path, ":")

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, "$ ")
		scanner.Scan()
		cmd := scanner.Text()

		err := RunCmd(cmd)

		if err != nil {
			fmt.Println(err)
			continue
		}
	}

}
