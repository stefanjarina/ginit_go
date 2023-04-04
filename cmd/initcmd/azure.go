package initcmd

import (
	"github.com/spf13/cobra"
	"github.com/stefanjarina/ginit/api"
)

var azureCmd = &cobra.Command{
	Use:   "azure [repo_name] [description]",
	Short: "Initialize repo for Azure DevOps",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		api.GetAnswers("azure")
	},
}

func init() {
}
