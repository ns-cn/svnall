package env

import "github.com/ns-cn/goter"

var (
	Branch   = goter.NewCmdFlagString(".", "branch", "b", "源分支")                                  // 仓库位置，不指定则为当前目录
	Authors  = goter.NewCmdFlagStringSlice([]string{}, "authors", "a", "通过提交者过滤（暂未生效）")           // 过滤用：作者
	Revision = goter.NewCmdFlagString("", "revision", "r", "版本，不指定则为所有，可选单次(1024)或范围(1024:2048)") // 修订版本区间，svn方式，可指定具体版本或版本区间
	Last     = goter.NewCmdFlagInt(0, "last", "l", "最近的多少次更新")                                    // 最近的修订版本次数指定，只查看最近的多少次提交
)

// merge使用的环境变量
var (
	Targets = goter.NewCmdFlagStringSlice([]string{}, "targets", "t", "目标分支")
	Preview = goter.NewCmdFlagBool(true, "preview", "p", "是否预览所有的变更，否则直接提交")
)
