package initcmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var githubCmd = &cobra.Command{
	Use:   "github [repo_name] [description]",
	Short: "Initialize repo for GitHub",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("github called")
	},
}

func init() {
}
