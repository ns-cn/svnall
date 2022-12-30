package action

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"svnall/env"
)

func SvnUpdateAll(repositories []env.Repository) error {
	for _, r := range repositories {
		SvnUpdPateRepository(r)
	}
	return nil
}

func SvnUpdPateRepository(repository env.Repository) {
	dir := repository.Dir
	if SvnUpdate(dir) && !env.FullThrough.Value {
		return
	}
	for _, subRepository := range repository.SubRepositories() {
		SvnUpdPateRepository(subRepository)
	}
}

func SvnUpdate(dir string) bool {
	stat, err := os.Stat(filepath.Join(dir, env.DIR_SVN))
	if err != nil || !stat.IsDir() {
		return false
	}
	if env.ShowURL.Value {
		information := Info(dir)
		fmt.Printf("svn update in:%s (%s)\n", dir, information.URL)
	} else {
		fmt.Printf("svn update in:%s\n", dir)
	}
	command := exec.Command("svn", "update")
	command.Dir = dir
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err = command.Run()
	return err == nil
}
