package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime"
	"strings"
)

// This package defines utility functions for some tasks

// DirFolder() returns the directory name of the directory
// you're in. The name is used in greating the project

func DirFolder() (string, error) {

	file, err := os.Getwd()
	if err != nil {
		log.Fatal("error reading the work directory path")
		return "", err
	}
	var seperatedPath []string
	if runtime.GOOS == "windows" {
		seperatedPath = strings.Split(file, string('\\'))
	} else {
		seperatedPath = strings.Split(file, string('/'))
	}

	// take the last value in the seperatedPath array,
	// that's the module name

	dir := seperatedPath[len(seperatedPath)-1]

	return dir, nil
}

// This is supplied a path and doesn't make use of the working directory
func DirPathFolder(path string) string {
	var seperatedPath []string
	if runtime.GOOS == "windows" {
		seperatedPath = strings.Split(path, string('\\'))
	} else {
		seperatedPath = strings.Split(path, string('/'))
	}

	dir := seperatedPath[len(seperatedPath)-1]

	return dir
}

// ListItems() list the items in the directory if any.

func ListItems() (string, error) {

	listError := errors.New("unable to list directory items, check if director exists")

	var resp string
	if runtime.GOOS == "windows" {
		resp = Powershell("ls") // from windows.go
		if resp == "" {
			return "", listError
		}
		return resp, nil
	} else {
		resp = Bash("ls") // from unix.go
		if resp == "" {
			return "", listError
		}
		return resp, nil
	}
}

// regCheck() is a regular expression function to check if a project
// has alrady being created in the target directory

func RegCheck(list string, keywords []string) (string, bool) {
	re := regexp.MustCompile(list)
	out := ""
	boolean := false
	for _, keyword := range keywords {
		if isFound := re.MatchString(keyword); isFound {
			out := fmt.Sprintf("%s exist", keyword)
			boolean = isFound
			return out, boolean
		}
	}
	return out, boolean
}
