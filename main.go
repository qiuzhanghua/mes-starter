package main

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/labstack/gommon/log"
	"os"
	"runtime"
)

func main() {

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Info("Alloc = ", byteToMB(m.Alloc))
	log.Info("TotalAlloc = ", byteToMB(m.TotalAlloc))
	log.Info("Sys = ", byteToMB(m.Sys))
	log.Info("NumGC = ", m.NumGC)

	privateKeyFile := "/Users/q/.ssh/id_rsa"
	publicKeys, err := ssh.NewPublicKeysFromFile("git", privateKeyFile, ""+
		"")
	if err != nil {
		log.Warnf("generate publickeys failed: %s\n", err.Error())
		return
	}
	url := "git@github.com:qiuzhanghua/learn-kotlin.git"
	// Clones the given repository in memory, creating the remote, the local
	// branches and fetching the objects, exactly as:
	log.Infof("git clone %s", url)

	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		Auth:     publicKeys,
		URL:      url,
		Progress: os.Stdout,
	})

	CheckIfError(err)

	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)
	log.Info("Alloc = ", byteToMB(m2.Alloc))
	log.Info("TotalAlloc = ", byteToMB(m2.TotalAlloc))
	log.Info("Sys = ", byteToMB(m2.Sys))
	log.Info("NumGC = ", m2.NumGC)

	// Gets the HEAD history from HEAD, just like this command:
	log.Info("git log")

	// ... retrieves the branch pointed by HEAD
	ref, err := r.Head()
	CheckIfError(err)

	// ... retrieves the commit history
	cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})
	CheckIfError(err)

	// ... just iterates over the commits, printing it
	err = cIter.ForEach(func(c *object.Commit) error {
		fmt.Println(c)
		return nil
	})
	CheckIfError(err)
}

func CheckIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func byteToMB(b uint64) uint64 {
	return b / 1024 / 1024
}
