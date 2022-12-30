package main

import (
	"fmt"
	"github.com/ns-cn/goter"
	"github.com/spf13/cobra"
	"os"
	"svnall/action"
	"svnall/env"
)

var cmdChanges = &goter.Command{
	Cmd: &cobra.Command{
		Use:     "changes",
		Aliases: []string{"list", "l"},
		Short:   "查询变更",
		Long: `功能： 【查询svn提交记录的文件变更】
必要参数：源分支branch(b)、版本范围（last(l)或revision(r)）
查看仓库最近10次提交： svnall changes -b ./code/svn -l 10
查看仓库指定版本号的提交： svnall changes -b ./code/svn -r 1024
查看仓库指定版本区间的提交： svnall changes -b ./code/svn -r 1024:2048
`,
		Run: func(cmd *cobra.Command, args []string) {
			fileUpdates, err := action.Log(env.Branch.Value, env.Authors.Value, env.Revision.Value, env.Last.Value)
			if err != nil {
				_, _ = os.Stderr.Write([]byte(err.Error()))
				_ = cmd.Help()
				return
			}
			for file, isDelete := range fileUpdates {
				if isDelete {
					fmt.Println("deleted\t" + file)
				} else {
					fmt.Println("updated\t" + file)
				}
			}
		},
	},
}
