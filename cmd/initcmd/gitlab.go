package initcmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var gitlabCmd = &cobra.Command{
	Use:   "gitlab [repo_name] [description]",
	Short: "Initialize repo for Gitlab",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gitlab called")
	},
}

func init() {
}
