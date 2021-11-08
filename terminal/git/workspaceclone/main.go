package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/src-d/go-git.v4/plumbing/transport"
)

const (
	// DefaultWorkspacePath is the default path to clone the workspace
	DefaultWorkspacePath = "/Users/bojun.cbj/workspace"
)

// workspace clone
func main() {
	argsWithoutProg := os.Args[1:]
	repoEndpoint := argsWithoutProg[0]
	endpoint, err := transport.NewEndpoint(repoEndpoint)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	localPath := endpoint.Path

	if strings.HasPrefix(endpoint.Path, "/") {
		localPath = endpoint.Path[1:]
	}

	if strings.HasSuffix(endpoint.Path, ".git") {
		localPath = endpoint.Path[:len(endpoint.Path)-4]
	}
	localPath = filepath.Join(DefaultWorkspacePath, endpoint.Host, localPath)
	fmt.Println(localPath)

	argsWithoutProg = append(argsWithoutProg, localPath)
	newCloneArgs := append([]string{"clone"}, argsWithoutProg...)
	fmt.Printf("real exec cmd: git %s\n", strings.Join(newCloneArgs, " "))

	err = exec.Command("git", newCloneArgs...).Run()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("cd %s\n", localPath)
	os.Exit(0)
}
