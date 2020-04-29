package cmd

import (
	"bytes"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"testing"
)

// executeCommand is a test utility function to execute the command
// with given args and returns the produced output and error
// as strings.
func executeCommand(root *cobra.Command, args ...string) (stdout string, stderr string, err error) {
	_, stdout, stderr, err = executeCommandC(root, args...)
	return stdout, stderr, err
}

// executeCommandC is a test utility function to execute the command
// with given args and returns the result command and the produced output and error
// as strings.
func executeCommandC(root *cobra.Command, args ...string) (c *cobra.Command, stdout string, stderr string, err error) {
	bufStdout := new(bytes.Buffer)
	bufStderr := new(bytes.Buffer)
	root.SetOut(bufStdout)
	root.SetErr(bufStderr)
	root.SetArgs(args)

	c, err = root.ExecuteC()

	return c, bufStdout.String(), bufStderr.String(), err
}

func TestRootVersion(t *testing.T) {
	root := GetRootCmd()
	stdout, stderr, err := executeCommand(root,"")
	t.Logf("Output:\n%v",stdout)
	t.Logf("Error:\n%v",stderr)
	assert.NoErrorf(t, err, "Command execution failed with error: %v", err)

}

var prCmdString = []string{"pr"}

func TestRoot_HasPrSubcommand(t *testing.T) {
	childCmd, _, err := GetRootCmd().Find(prCmdString)
	assert.NoErrorf(t, err, "Searching for sub-command %v failed with error: %v", prCmdString[0], err)
	assert.Equalf(t, "pr", childCmd.Name(), "We didn't find a 'pr' sub-command")
}

func TestRoot_PrSubcommandRunsWithoutError(t *testing.T) {
	stdout, stderr, err := executeCommand(GetRootCmd(), prCmdString...)
	t.Logf("Output:\n%v",stdout)
	t.Logf("Error:\n%v",stderr)
	assert.NoErrorf(t, err, "Command execution failed with error: %v", err)
}


