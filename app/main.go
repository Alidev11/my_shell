package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main() {
	// TODO: Uncomment the code below to pass the first stage
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		// Print $ and read commands until exiting the loop
		fmt.Fprint(os.Stdout, "$ ")
		scanner.Scan()
		cmd := scanner.Text()
		
		// Exit command
		if strings.HasPrefix(cmd, "exit") {
			parts := strings.Fields(cmd)
			if(len(parts) == 2) {
				i, err := strconv.Atoi(parts[1])
				if err != nil {
					fmt.Println("Status code is wrong: ", err)
					continue
				} else {
					os.Exit(i)
				}
			}
		} else {
			fmt.Fprint(os.Stdout, cmd + ": command not found\n")
		}
	}
	

}
