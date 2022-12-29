package main

import (
	"github.com/ns-cn/goter"
	"github.com/spf13/cobra"
	"os"
	"svnall/action"
	"svnall/env"
)

var cmdUpdate = &goter.Command{
	&cobra.Command{
		Use:   "update",
		Short: "tool to update multi svn repositories",
		Run: func(cmd *cobra.Command, args []string) {
			repositories, err := env.InitEnv(args)
			if err != nil {
				_, _ = os.Stderr.WriteString(err.Error())
				_ = cmd.Help()
				return
			}
			_ = action.SvnUpdateAll(repositories)
		},
	},
}
