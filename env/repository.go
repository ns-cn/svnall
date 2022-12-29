package env

import (
	"os"
	"path/filepath"
)

type Repository struct {
	Dir   string // 目录位置
	Depth int    // 查询深度
}

func (r Repository) SubRepositories() []Repository {
	subDirs, _ := os.ReadDir(r.Dir)
	subRepositories := make([]Repository, 0)
	for _, subDir := range subDirs {
		if subDir == nil || !subDir.IsDir() || subDir.Name() == DIR_SVN {
			continue
		}
		subRepository := r.stepIn(subDir.Name())
		if subRepository != nil {
			subRepositories = append(subRepositories, *subRepository)
		}
	}
	return subRepositories
}

func (r Repository) stepIn(subDir string) *Repository {
	if Depth.Value <= 0 {
		return nil
	}
	return &Repository{Dir: filepath.Join(r.Dir, subDir), Depth: r.Depth - 1}
}
