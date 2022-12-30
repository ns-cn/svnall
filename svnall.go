package main

import (
	"github.com/ns-cn/goter"
	"svnall/env"
)

func main() {
	root := goter.NewRootCmd("svnall", "tool to update multi svn repository update ", env.VERSION)
	root.AddCommand(cmdUpdate.Bind(&env.Depth, &env.Exclude, &env.FullThrough, &env.ShowURL))
	root.AddCommand(cmdChanges.Bind(&env.Authors, &env.Branch, &env.Revision, &env.Last))
	root.AddCommand(cmdMerge.Bind(&env.Authors, &env.Branch, &env.Revision, &env.Last, &env.Targets, &env.Preview))
	// 数据源
	_ = root.Execute()
}
