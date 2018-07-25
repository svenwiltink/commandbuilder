package commandbuilder

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os/exec"
)

func TestCommandBuilder_New(t *testing.T) {
	builder := New("test")

	assert.Equal(t, builder.Prefix, "test")
	assert.Empty(t, builder.Arguments)
}

func TestCommandBuilder_AddArgument(t *testing.T) {
	builder := New("test").AddArgument("test")

	assert.Equal(t, builder.Prefix, "test")
	assert.Equal(t, builder.Arguments, []string{"test"})
}

func TestCommandBuilder_AddArgument_chained(t *testing.T) {
	builder := New("test").
		AddArgument("test").
		AddArgument("test2")

	assert.Equal(t, builder.Prefix, "test")
	assert.Equal(t, builder.Arguments, []string{"test", "test2"})
}

func TestCommandBuilder_AddArguments(t *testing.T) {
	builder := New("test").AddArguments("test")

	assert.Equal(t, builder.Prefix, "test")
	assert.Equal(t, builder.Arguments, []string{"test"})
}

func TestCommandBuilder_AddArguments_multiple(t *testing.T) {
	builder := New("test").AddArguments("test", "test2")

	assert.Equal(t, builder.Prefix, "test")
	assert.Equal(t, builder.Arguments, []string{"test", "test2"})
}

func TestCommandBuilder_Tobuilder(t *testing.T) {
	builder := New("test").AddArguments("test", "test2")
	cmd := builder.ToCmd()

	assert.Equal(t, cmd, exec.Command("test", "test", "test2"))
}