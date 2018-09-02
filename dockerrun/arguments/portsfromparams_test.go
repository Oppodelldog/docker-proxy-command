package arguments

import (
	"testing"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockerrun/builder/mocks"
	"github.com/stretchr/testify/assert"
	"os"
)

func TestBuildPortsFromParams_portsDefined_matchWithArgs(t *testing.T) {
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	os.Args = []string{"-test1", "--inspect-brk=78129", "-colors=on"}

	portsFromParams := []string{
		"--inspect-brk=(\\d*)",
	}
	commandDef := &config.CommandDefinition{
		PortsFromParams: &portsFromParams,
	}

	builder := &mocks.Builder{}

	builder.On("AddPortMapping", "78129:78129").Return(builder)

	BuildPortsFromParams(commandDef, builder)

	builder.AssertExpectations(t)
}

func TestBuildPortsFromParams_portsDefined_2matchesWithArgs(t *testing.T) {
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	os.Args = []string{"-test1", "--inspect-brk=78129", "-colors=on","--inspect=2938"}

	portsFromParams := []string{
		"--inspect-brk=(\\d*)",
		"--inspect=(\\d*)",
	}
	commandDef := &config.CommandDefinition{
		PortsFromParams: &portsFromParams,
	}

	builder := &mocks.Builder{}

	builder.On("AddPortMapping", "78129:78129").Return(builder)
	builder.On("AddPortMapping", "2938:2938").Return(builder)

	BuildPortsFromParams(commandDef, builder)

	builder.AssertExpectations(t)
}

func TestBuildPortsFromParams_portsDefined_noMatchWithArgs(t *testing.T) {
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	os.Args = []string{"-test1", "-colors=on"}

	portsFromParams := []string{
		"--inspect-brk=(\\d*)",
	}
	commandDef := &config.CommandDefinition{
		PortsFromParams: &portsFromParams,
	}

	builder := &mocks.Builder{}

	BuildPortsFromParams(commandDef, builder)

	builder.AssertNumberOfCalls(t,"AddPortMapping",0)
}

func TestBuildPortsFromParams_portsNotDefined(t *testing.T) {
	commandDef := &config.CommandDefinition{
		PortsFromParams: nil,
	}

	builder := &mocks.Builder{}

	BuildPortsFromParams(commandDef, builder)

	assert.Empty(t, builder.Calls)
}
