/*
Copyright © 2020 Flanksource

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/flanksource/gohub-helper/pkg/pr"
	"github.com/spf13/cobra"
)



// prCmd represents the pr command
var prCmd = cobra.Command{
	Use:   "pr",
	Short: "manipulate PRs",
	Long: `PR is identified with:
-n [PR #] -u [User] -r [Repository]`,
	//RunE: func(cmd *cobra.Command, args []string) error {
	//	fmt.Println("pr called")
	//	return nil
	//},
}

// prCommentCmd represents the comment command
var prCommentCmd = cobra.Command{
	Use:   "comment",
	Short: "manipulate PR issue comments",
	Long: `PR is identified with:
-n [PR #] -u [User] -r [Repository]`,
	//RunE: func(cmd *cobra.Command, args []string) error {
	//	fmt.Println("comment called")
	//	return nil
	//},
}

// prCommentPostCmd represents the comment post command
var prCommentPostCmd = cobra.Command{
	Use:   "post",
	Short: "post a new PR issue comment",
	Long: `PR is identified with:
-n [PR #] -u [User] -r [Repository]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var failed bool = false
		var failMsg string = ""
		findStringFlagOrFail := func (name string) string {
			value, err := cmd.Flags().GetString(name)
			if err != nil {
				failed = true
				failMsg += fmt.Sprintf("Couldn't parse %v value", name)
			}
			return value
		}
		findIntFlagOrFail := func (name string) int {
			value, err := cmd.Flags().GetInt(name)
			if err != nil {
				failed = true
				failMsg += fmt.Sprintf("Couldn't parse %v value\n", name)
			}
			return value
		}
		comment := pr.PRIssueComment{
			PR:  pr.PR{
				Owner: findStringFlagOrFail("owner"),
				Repo:  findStringFlagOrFail("repo"),
				Num:   findIntFlagOrFail("num"),
			},
			Msg: findStringFlagOrFail("msg"),
		}
		if failed {
			return fmt.Errorf(failMsg)
		}
		err := comment.Post()
		if err!= nil {
			cmd.PrintErrf("Error: %v",err)
			return err
		}
		return nil
	},

}

// GetPrCmd gives an independent prCmd Command copy
func GetPrCmd() *cobra.Command {
	pr := prCmd
	pr.PersistentFlags().IntP("num",  "n",0,"The PR number")
	pr.PersistentFlags().StringP("owner", "o", "", "The owner user/organisation")
	pr.PersistentFlags().StringP("repo", "r", "", "The repo")

	comment := prCommentCmd
	post := prCommentPostCmd
	post.Flags().StringP( "msg", "m", "", "The PR Issue Comment Message")

	comment.AddCommand(&post)
	pr.AddCommand(&comment)
	return &pr
}

func init() {

	//prCmd.PersistentFlags().IntVarP(&_PR.Num,"num", "n", 0,"The PR number")
	//prCmd.PersistentFlags().StringVarP(&_PR.Owner, "owner", "o", "", "The owner user/organisation")
	//prCmd.PersistentFlags().StringVarP(&_PR.Repo, "repo", "r", "", "The repo")

	//prCommentPostCmd.Flags().StringVarP(&Comment.Msg, "message", "m", "", "The PR Issue Comment Message")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// prCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// prCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
