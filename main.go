package main

import (
	"fmt"
	"log"
	"os"
)

var cacheLocation string
var arg []string

func addJumpPoint(pref string) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error in fetching the current directory: %v\n", err)
	}
	log.Printf("adding %s to cache\n", dir)
	configFilePath := CreateConfig(cacheLocation)
	AddKeyToStore(configFilePath, Jumpo{
		Location: dir,
		Prefix:   pref,
	})
}

func removeJumpPoint(name string) {
	RemoveKey(ConfigFileLocation(cacheLocation), name)
}

func listJumpPoint() {
	ListData(ConfigFileLocation(cacheLocation))
}

func jumpPoint(name string) {
	if val := FetchValue(ConfigFileLocation(cacheLocation), name); val != "" {
		// Since each executable is running in its own forked process
		// and cannot change the directory of parent shell so need to output
		// the valid shell command on the stdout which then can be used in backticks
		// or explicit execution from shell.
		fmt.Printf("cd %s", val)
	}
}

func main() {
	arg = os.Args[1:]

	if len(arg) < 1 {
		os.Exit(0)
	}

	if cacheDir := GetCacheDir(); cacheDir != "" {
		cacheLocation = cacheDir
		if err := createIfNotExists(cacheDir); err == nil {
			log.Printf("Cache directory created successfully\n")
		}
	}

	switch arg[0] {
	case "-add", "-a":
		if len(arg) < 2 {
			log.Fatalf("Require short-hand as second arg\n")
		}
		addJumpPoint(arg[1])
	case "-remove", "-r":
		if len(arg) < 2 {
			log.Fatalf("Require short-hand as second arg\n")
		}
		removeJumpPoint(arg[1])
	case "-list", "-l":
		listJumpPoint()
	default:
		jumpPoint(arg[0])
	}
}
