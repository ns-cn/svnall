# svnall
terminal tools to operate svn for all svn repositories

## svnall update
用于更新目标svn仓库，可选择环境变量`SVNALL_REPOSITORIES`配置仓库，或命令行参数传递多个svn仓库
```shell
Usage:
  svnall update [flags] [repository...]

Flags:
  -d, --Depth int   the Depth searching in dir(不指定则使用环境变量：SVNALL_DEPTH，默认值2) (default -1)
  -e, --exclude     是否排除环境变量配置仓库,默认不排除（环境变量：SVNALL_REPOSITORIES）
  -h, --help        help for update
  -u, --showurl     是否显示仓库的完整URL (default true)
  -t, --through     是否已经找到.svn继续往下查找，默认不继续往下
```
其中`repository`的配置格式为`{仓库文件夹路径}[#查询深度]`,`查询深度`可选

例如```~/workspace/code1 ~/workspace/code2#4```

## svnall changes
用于按照一定的查询条件列举指定仓库的提交历史
```shell
Usage:
  svnall changes [flags]

Aliases:
  changes, list, l

Flags:
  -a, --authors strings   通过提交者过滤（暂未生效）
  -b, --branch string     源分支 (default ".")
  -h, --help              help for changes
  -l, --last int          最近的多少次更新
  -r, --revision string   版本，不指定则为所有，可选单次(1024)或范围(1024:2048)
```

## svnall merge
用于按照一定的查询条件将源仓库的变更合并到其他的多个仓库中
```shell
Usage:
  svnall merge [flags]

Aliases:
  merge, m

Flags:
  -a, --authors strings   通过提交者过滤（暂未生效）
  -b, --branch string     源分支 (default ".")
  -h, --help              help for merge
  -l, --last int          最近的多少次更新
  -p, --preview           是否预览所有的变更，否则直接提交 (default true)
  -r, --revision string   版本，不指定则为所有，可选单次(1024)或范围(1024:2048)
  -t, --targets strings   目标分支
```