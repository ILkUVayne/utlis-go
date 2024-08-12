package flie

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
	"sync"
)

var (
	rw sync.RWMutex
)

// Home returns the home directory for the executing user.
//
// This uses an OS-specific method for discovering the home directory.
// An error is returned if a home directory cannot be detected.
//
// e.g. return "~/" dir such as "/home/tom" in linux
func Home() (string, error) {
	usr, err := user.Current()
	if nil == err {
		return usr.HomeDir, nil
	}
	// cross compile support
	if "windows" == runtime.GOOS {
		return homeWindows()
	}
	// Unix-like system, so just assume Unix
	return homeUnix()
}

func homeUnix() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}
	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}
	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}
	return result, nil
}

func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}
	return home, nil
}

func IsExists(path string) (os.FileInfo, bool) {
	rw.RLock()
	defer rw.RUnlock()
	f, err := os.Stat(path)
	return f, err == nil || !errors.Is(err, os.ErrNotExist)
}

func IsDir(path string) (os.FileInfo, bool) {
	rw.RLock()
	defer rw.RUnlock()
	f, flag := IsExists(path)
	return f, flag && f.IsDir()
}

func IsFile(path string) (os.FileInfo, bool) {
	rw.RLock()
	defer rw.RUnlock()
	f, flag := IsExists(path)
	return f, flag && !f.IsDir()
}
