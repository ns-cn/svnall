package main

import (
	"github.com/ns-cn/goter"
	"svnall/env"
)

func main() {
	root := goter.NewRootCmd("svnall", "tool to update multi svn repository update ", env.VERSION)
	root.AddCommand(cmdUpdate.Bind(env.Depth, env.Exclude, env.FullThrough))
	// 数据源
	_ = root.Execute()
}
