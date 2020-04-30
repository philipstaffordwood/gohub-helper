package test

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"testing"
)

// ExecuteCommand is a test utility function to execute the command
// with given args and returns the produced output and error
// as strings.
func ExecuteCommand(root *cobra.Command, args ...string) (stdout string, stderr string, err error) {
	_, stdout, stderr, err = ExecuteCommandC(root, args...)
	return stdout, stderr, err
}

// ExecuteCommandC is a test utility function to execute the command
// with given args and returns the result command and the produced output and error
// as strings.
func ExecuteCommandC(root *cobra.Command, args ...string) (c *cobra.Command, stdout string, stderr string, err error) {
	bufStdout := new(bytes.Buffer)
	bufStderr := new(bytes.Buffer)
	root.SetOut(bufStdout)
	root.SetErr(bufStderr)
	root.SetArgs(args)

	c, err = root.ExecuteC()

	return c, bufStdout.String(), bufStderr.String(), err
}
// HasSubcommand is an assertion helper that verifies if cobra.Command cmd has a specific
// child command
func HasSubcommand(t *testing.T, cmd *cobra.Command, name string, msgAndArgs ...interface{} ) bool {
	targetCmd, _, err := cmd.Find([]string {name})
	if err != nil || name != targetCmd.Name() {
		return assert.Fail(t, fmt.Sprintf("Command '%v', expected subcommand '%v'",cmd.Name(),name), msgAndArgs...)
	}
	return true
}
