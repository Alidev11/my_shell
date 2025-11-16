package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"os/exec"
)

var CmdList = [10]string{"echo", "exit", "type"}

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
		if contains(CmdList, arg) {
			fmt.Println(arg, "is a shell builtin")
		} else {
			res, returnedPath := FileExists(Arr, arg)
			
			if res == true {
				fmt.Println(arg, "is", returnedPath)
			} else {
				return fmt.Errorf("%s: not found", arg)
			}
		}
	} else {
		return fmt.Errorf("%s: not found", cmd)
	}

	return nil
}

func FileExists(dirs []string, target string) (bool, string) {
	returnedPath := ""
	// loop through directories
	for _, dir := range dirs {
		// fmt.Println("Dir: " , dir)
		files, err := os.ReadDir(dir)
		if err != nil {
			continue
		}

		// loop through files
		for _, file := range files {
			info, err := file.Info()
			if err != nil {
				continue
			}
			perm := info.Mode().Perm().String()
			// check only execution permission for either(owner, group, other)
			fileName := strings.Split(file.Name(), ".")
			if (fileName[0] == target) && (perm[3] == 'x' || perm[6] == 'x' || perm[9] == 'x') {
				returnedPath = dir + "/" + target
				return true, returnedPath
			} else {
				continue
			}
		}
	}
	return false, returnedPath
}

// ------------------------
func RunCmd(cmd string) error {
	cmdArr := strings.Split(cmd, " ")
	if strings.HasPrefix(cmd, "exit") {
		err := Exit(cmd)
		return err
	} else if strings.HasPrefix(cmd, "echo") {
		err := Echo(cmd)
		return err
	} else if strings.HasPrefix(cmd, "type") {
		err := Type(cmd)
		return err
	} else if res, _ := FileExists(Arr, cmdArr[0]); res == true{
		cmd := exec.Command(cmdArr[0], cmdArr[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}	else {
		fmt.Fprint(os.Stdout, cmd+": command not found\n")
	}
	return nil
}
