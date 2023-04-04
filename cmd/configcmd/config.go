package configcmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stefanjarina/ginit/globals"
	"github.com/stefanjarina/ginit/utils"
	"strings"
)

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func addSubcommands() {
	ConfigCmd.AddCommand(allCmd)
	ConfigCmd.AddCommand(getCmd)
	ConfigCmd.AddCommand(setCmd)
	ConfigCmd.AddCommand(removeCmd)
}

func init() {
	repos := utils.NewEnum(globals.SupportedRepos, "")
	ConfigCmd.PersistentFlags().VarP(repos, "repo", "r", fmt.Sprintf("Specify repository <%s>", strings.Join(repos.Allowed, "|")))
	if err := ConfigCmd.MarkPersistentFlagRequired("repo"); err != nil {
		fmt.Println(err)
	}
	
	addSubcommands()
}
