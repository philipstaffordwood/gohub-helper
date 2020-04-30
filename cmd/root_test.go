package cmd

import (
	"github.com/flanksource/gohub-helper/cmd/test"
	"github.com/stretchr/testify/assert"
	"testing"
)



func TestRootVersion(t *testing.T) {
	root := GetRootCmd()
	stdout, stderr, err := test.ExecuteCommand(root,"")
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
	stdout, stderr, err := test.ExecuteCommand(GetRootCmd(), prCmdString...)
	t.Logf("Output:\n%v",stdout)
	t.Logf("Error:\n%v",stderr)
	assert.NoErrorf(t, err, "Command execution failed with error: %v", err)
}


