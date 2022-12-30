package env

import (
	"fmt"
	"github.com/ns-cn/goter"
	"os"
	"strconv"
	"strings"
)

const (
	ENV_REPOSITORIES   = "SVNALL_REPOSITORIES"
	ENV_DEPTH          = "SVNALL_DEPTH"
	DEFAULT_DEPTH      = 2 // 默认的遍历深度
	SPLITOR_DEPTH      = "#"
	SPLITOR_REPOSITORY = ";"
	NO_ERR             = 0 // 无异常
	ERR_FORMAT_ARG     = 1 // 参数格式
	ERR_ACCESS         = 2 // 路径不合法,或无访问权限

	DIR_SVN = ".svn" // svn目录
)

var (
	Depth       = goter.NewCmdFlagInt(-1, "Depth", "d", "the Depth searching in dir(不指定则使用环境变量：SVNALL_DEPTH，默认值2)")
	Exclude     = goter.NewCmdFlagBool(false, "exclude", "e", "是否排除环境变量配置仓库,默认不排除（环境变量：SVNALL_REPOSITORIES）")
	FullThrough = goter.NewCmdFlagBool(false, "through", "t", "是否已经找到.svn继续往下查找，默认不继续往下")
	ShowURL     = goter.NewCmdFlagBool(true, "showurl", "u", "是否显示仓库的完整URL")
)

func InitEnv(args []string) (repositories []Repository, err error) {
	// no user input or wrong user input
	if Depth.Value < 0 {
		envDepth := os.Getenv(ENV_DEPTH)
		parsedDepth, err := strconv.Atoi(envDepth)
		if err != nil || parsedDepth < 0 {
			Depth.Value = DEFAULT_DEPTH
		} else {
			Depth.Value = parsedDepth
		}
	}
	repositories = make([]Repository, 0)
	wrongRepositories := make([]string, 0)
	if err != nil {
		return
	}
	unparsedRepositories := make([]string, 0)
	if len(args) > 0 {
		unparsedRepositories = append(unparsedRepositories, args...)
	}
	if !Exclude.Value {
		envRepositories := os.Getenv(ENV_REPOSITORIES)
		for _, unparsedRepository := range strings.Split(envRepositories, SPLITOR_REPOSITORY) {
			if unparsedRepository != "" {
				unparsedRepositories = append(unparsedRepositories, unparsedRepository)
			}
		}
	}
	if unparsedRepositories == nil || len(unparsedRepositories) == 0 {
		err = fmt.Errorf("尚未指定更新仓库地址,可选参数或环境变量方式\n")
		return
	}
	// 从参数中读取
	for _, unparsedRepository := range unparsedRepositories {
		repository, err := ParseRepository(unparsedRepository)
		if err > 0 {
			wrongRepositories = append(wrongRepositories, unparsedRepository)
		} else {
			repositories = append(repositories, repository)
		}
	}
	if len(wrongRepositories) != 0 {
		err = fmt.Errorf("wrong repository: %v\n", wrongRepositories)
	}
	return
}

/*
arg2Repo: 将特定的参数转换为对应的仓库地址信息
格式要求:仓库地址路径[#寻址深度]
*/
func ParseRepository(unparsedRepository string) (Repository, int) {
	splited := strings.Split(unparsedRepository, SPLITOR_DEPTH)
	if len(splited) > 2 || splited[0] == "" {
		return Repository{}, ERR_FORMAT_ARG
	} else {
		var simpleDepth = Depth.Value
		var err error
		if len(splited) == 2 && splited[1] != "" {
			simpleDepth, err = strconv.Atoi(splited[1])
			if err != nil {
				return Repository{}, ERR_FORMAT_ARG
			}
		}
		dir := splited[0]
		if !isDirExists(dir) {
			return Repository{}, ERR_ACCESS
		}
		return Repository{Dir: dir, Depth: simpleDepth}, NO_ERR
	}
}

func isDirExists(path string) bool {
	state, err := os.Stat(path)
	return err == nil && state.IsDir()
}
