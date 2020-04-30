package cmd

import (
	"github.com/flanksource/gohub-helper/cmd/test"
	"github.com/stretchr/testify/assert"
	"strings"

	"testing"
)

var prConnectString = "pr -n 8 -o philipstaffordwood -r hello-go-githubapp"


func TestPr_CanSpecifyPrWithFlags(t *testing.T) {
	args := strings.Split(prConnectString, " ")
	cmd := GetPrCmd()
	cmd.ParseFlags(args)
	fn := cmd.Flag("num")
	assert.Equalf(t, "8", fn.Value.String(), "Not the PR number we expected in flag")


	fown := cmd.Flag("owner")
	assert.Equalf(t, "philipstaffordwood", fown.Value.String(), "Not the owner we expected in flag")


	frepo:= cmd.Flag("repo")
	assert.Equalf(t, "hello-go-githubapp", frepo.Value.String(), "Not the owner we expected in flag")


}


func TestPr_HasCommentSubcommand(t *testing.T) {
	test.HasSubcommand(t, GetPrCmd(), "comment", "We didn't find a 'comment' sub-command", )
}

func TestPrComment_HasPostSubcommand(t *testing.T) {
	cmd, _, err := GetPrCmd().Find([]string{"comment"})
	assert.NoErrorf(t, err, "Error finding '%v' subcommand comment.post : %v", err)
	test.HasSubcommand(t, cmd, "comment", "We didn't find a 'comment' sub-command", )
}

var prCommentPostString = "pr comment post -m 'hello' -n 8 -o philipstaffordwood -r hello-go-githubapp"

func TestPrCommentPost_CanSpecifyMessageWithFlags(t *testing.T) {
	args := strings.Split(prCommentPostString, " ")
	post := []string{"comment","post"}
	cmdPr := GetPrCmd()

	cmdPost, _, err := cmdPr.Find(post)
	assert.NoErrorf(t, err, "Searching for comment post sub-command %v failed with error: %v", post[0], err)
	err = cmdPost.ParseFlags(args[2:])
	assert.NoErrorf(t, err, "Parsing flags failed with error: %v", err)
	fm := cmdPost.Flag("msg")
	assert.NotNil(t, fm, "msg flag not found")
	assert.Equalf(t, "'hello'", fm.Value.String(), "Not the message we expected in flag")
}