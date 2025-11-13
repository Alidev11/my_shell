package main

import (
	"bufio"
	"fmt"
	"os"
)

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
