package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var CmdList = [10]string{"echo", "exit", "type", "pwd", "cd"}

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
			res, returnedPath := FileExists(DIRS, arg)

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

func Cd(cmd string) error{
	parts := strings.Fields(cmd)
	if len(parts)==2{
		arg := parts[1]
		if strings.HasPrefix(arg, "/") {
			_, err := os.Stat(arg);
			if err != nil && os.IsNotExist(err){
				return fmt.Errorf("%s%s%s", "cd: ",arg,": No such file or directory");
			}
			os.Chdir(arg);
		}
	}else{
		return fmt.Errorf("%s", "cd command with 0 arguments!");
	}

	return nil;
}

func FileExists(dirs []string, target string) (exists bool, path string) {
	returnedPath := ""
	for _, dir := range dirs {
		file, err := os.Stat(dir + "/" + target)

		if err == nil {
			perm := file.Mode().Perm().String()
			if perm[3] == 'x' {
				returnedPath = dir + "/" + target
				return true, returnedPath
			}
		}
	}
	return false, returnedPath
}

// ------------------------
func RunCmd(cmd string) error {
	cmdArr := strings.Split(cmd, " ")
	switch cmdArr[0] {
	case "exit":
		err := Exit(cmd)
		return err
	case "echo":
		err := Echo(cmd)
		return err
	case "type":
		err := Type(cmd)
		return err
	case "pwd":
		wd, err := os.Getwd()
		fmt.Println(wd)
		return err
	case "cd":
		err := Cd(cmd);
		return err;
	default:
		if res, _ := FileExists(DIRS, cmdArr[0]); res == true {
			cmd := exec.Command(cmdArr[0], cmdArr[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
		} else {
			fmt.Fprint(os.Stdout, cmd+": command not found\n")
		}
	}
	return nil
}
