package initcmd

import (
	"github.com/spf13/cobra"
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize repo",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func addSubcommands() {
	InitCmd.AddCommand(azureCmd)
	InitCmd.AddCommand(githubCmd)
	InitCmd.AddCommand(gitlabCmd)
}

func init() {
	addSubcommands()
}
