package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cmds = [10]string{"echo", "exit", "type"}

func Exit(cmd string) error {
	parts := strings.Fields(cmd)
	if len(parts) == 2 {
		i, err := strconv.Atoi(parts[1])
		if err != nil {
			return fmt.Errorf("Status code is wrong: %s", err)
		} else {
			os.Exit(i)
		}
	}
	return nil
}

func Echo(cmd string) error {
	str, _ := strings.CutPrefix(cmd, "echo ")
	parts := strings.Fields(cmd)

	if len(parts) >= 2 {
		fmt.Println(str)
	} else {
		return fmt.Errorf("Echo command with no argument!")
	}

	return nil
}

func contains(slice [10]string, val string) bool {
	for _, s := range slice {
		if s == val {
			return true
		}
	}
	return false
}

func Type(cmd string) error {
	parts := strings.Fields(cmd)
	if len(parts) == 2 {
		arg := parts[1]
		if contains(cmds, arg) {
			fmt.Println(arg, "is a shell builtin")
		} else {
			return fmt.Errorf("%s: not found", arg)
		}
	}else {
		return fmt.Errorf("%s: not found", cmd)
	}

	return nil
}

// ------------------------
func RunCmd(cmd string) error {
	if strings.HasPrefix(cmd, "exit") {
		err := Exit(cmd)
		return err
	} else if strings.HasPrefix(cmd, "echo") {
		err := Echo(cmd)
		return err
	} else if strings.HasPrefix(cmd, "type") {
		err := Type(cmd)
		return err
	} else {
		fmt.Fprint(os.Stdout, cmd+": command not found\n")
	}

	return nil
}
