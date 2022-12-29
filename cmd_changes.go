package main

import (
	"bufio"
	"fmt"
	"github.com/ns-cn/goter"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
	"strings"
	"svnall/action"
	"svnall/env"
)

var cmdChanges = &goter.Command{
	&cobra.Command{
		Use:     "changes",
		Aliases: []string{"l"},
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

var cmdMerge = &cobra.Command{
	Use:     "merge",
	Aliases: []string{"m"},
	Short:   "合并变更",
	Long: `功能： 【查询svn提交记录的文件变更并合并到指定目录】
必要参数：源分支branch(b)、目标分支target(t,可多选)、版本范围（last(l)或revision(r)）
存放仓库最近10次提交到指定目录： muddler merge -b ./code/svn -t ./update-20220818 -l 10
合并指定分支版本到另外多个分支： muddler merge -b ./code/svn -t ./branch1 -t ./branch2 -r 1024
合并仓库指定版本区间的提交到另外一个分支： muddler merge -b ./code/svn -r 1024:2048 -t ./branch1
可选增加preview用于是否预览，还是直接合并
`,
	Run: func(cmd *cobra.Command, args []string) {
		workDir, err := os.Getwd()
		if err != nil {
			_, _ = os.Stderr.WriteString("请在命令行中运行！\n")
		}
		sourceBranch, err := os.Stat(env.Branch.Value)
		if err != nil {
			_, _ = os.Stderr.WriteString(err.Error() + "\n")
			return
		}
		fmt.Println(sourceBranch, workDir)
		if len(env.Targets.Value) == 0 {
			_, _ = os.Stderr.WriteString("未指定输出目录，参考参数--target\n")
			_ = cmd.Help()
			return
		}

		fileUpdates, err := action.Log(env.Branch.Value, env.Authors.Value, env.Revision.Value, env.Last.Value)
		if err != nil && err != io.EOF {
			_, _ = os.Stderr.WriteString(err.Error())
			return
		}
		_ = os.Chdir(workDir)
		if len(fileUpdates) == 0 {
			_, _ = os.Stderr.WriteString(fmt.Sprintf("在版本[%s]未发现变更\n", env.Revision))
			return
		}
		merge := false
		if env.Preview.Value {
			_, _ = os.Stdout.WriteString("[预览]包含如下变更：\n")
			for file, isDelete := range fileUpdates {
				if isDelete {
					fmt.Println("deleted\t" + file)
				} else {
					fmt.Println("updated\t" + file)
				}
			}
			inputReader := bufio.NewReader(os.Stdin)
			for {
				_, _ = os.Stdout.WriteString("是否进行合并(Y/N):")
				userInput, _ := inputReader.ReadString('\n')
				if strings.ToUpper(userInput)[:1] == "Y" {
					_, _ = os.Stdout.WriteString("进行合并！\n")
					merge = true
					break
				} else if strings.ToUpper(userInput)[:1] == "N" {
					_, _ = os.Stdout.WriteString("取消合并！\n")
					break
				}
			}
		}
		if merge {
			for _, target := range env.Targets.Value {
				_, _ = os.Stdout.WriteString("→：" + target + "\n")
				target, err = filepath.Abs(target)
				if err != nil {
					fmt.Println(err)
					continue
				}
				for file, isDelete := range fileUpdates {
					targetFile := action.GetFilePath(target, file)
					if isDelete {
						fmt.Println("deleted\t" + targetFile)
					} else {
						sourceFile := action.GetFilePath(env.Branch.Value, file)
						stat, err := os.Stat(sourceFile)
						if err != nil {
							_, _ = os.Stderr.WriteString("updated error: " + err.Error() + "\n")
							continue
						}
						if stat.IsDir() {
							_, _ = os.Stderr.WriteString(fmt.Sprintf("update skiped dir: %s\n", sourceFile))
							continue
						}
						err = action.CopyFile(sourceFile, targetFile)
						if err != nil {
							_, _ = os.Stderr.WriteString("updated error: " + err.Error() + "\n")
						}
					}
				}
			}
		}
	},
}
