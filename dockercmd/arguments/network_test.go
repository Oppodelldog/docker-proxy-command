package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercmd/builder/mocks"
	"github.com/testify/assert"
	"testing"
)

func TestBuildNetwork_NetworkIsTrue(t *testing.T) {
	network := "my-docker-network"
	commandDef := &config.CommandDefinition{
		Network: &network,
	}
	builder := &mocks.Builder{}

	builder.On("SetNetwork", network).Return(builder)

	BuildNetwork(commandDef, builder)

	builder.AssertExpectations(t)
}

func TestBuildNetwork_NetworkIsFalse(t *testing.T) {
	commandDef := &config.CommandDefinition{}
	builder := &mocks.Builder{}

	BuildNetwork(commandDef, builder)

	assert.Empty(t, builder.Calls)
}
