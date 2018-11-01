package arguments

import (
	"reflect"
	"testing"

	"fmt"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBuildName_NameIsSet_AndNotUnique_ExpectAppropriateBuilderCall(t *testing.T) {
	randomNamePartStub := "123"

	testCases := []struct {
		testNo                int
		containerName         string
		uniqueNames           bool
		expectedContainerName string
	}{
		{testNo: 1, containerName: "my-container", uniqueNames: false, expectedContainerName: "my-container"},
		{testNo: 2, containerName: "my-container", uniqueNames: true, expectedContainerName: "my-container123"},
	}

	for _, testCase := range testCases {

		t.Logf("testCase no: %v", testCase.testNo)

		containerName := testCase.containerName
		uniqueNames := testCase.uniqueNames
		commandDef := &config.CommandDefinition{
			Name:        &containerName,
			UniqueNames: &uniqueNames,
		}
		builder := &mocks.Builder{}
		builder.On("SetContainerName", testCase.expectedContainerName).Return(builder)

		nameBuilder := nameArgumentBuilder{nameRandomizerFunc: func(containerName string) string { return fmt.Sprintf("%s%v", containerName, randomNamePartStub) }}
		nameBuilder.BuildArgument(commandDef, builder)

		builder.AssertExpectations(t)
	}
}

func TestBuildName_NameIsNotSet_AndNotUniqueNames(t *testing.T) {
	commandDef := &config.CommandDefinition{}
	builder := &mocks.Builder{}

	nameBuilder := NewNameArgumentBuilder()
	nameBuilder.BuildArgument(commandDef, builder)

	assert.Empty(t, builder.Calls)
}

func TestBuildName_ArgumentsBuilderNameRandomizerFunc(t *testing.T) {
	if reflect.ValueOf(NewNameArgumentBuilder().(*nameArgumentBuilder).nameRandomizerFunc).Pointer() != reflect.ValueOf(defaultNameRandomizerFunc).Pointer() {
		t.Fatalf("nameArgumentBuilder.nameRandomizerFunc is not set to defaultNameRandomizerFunc")
	}
}

func Test_defaultNameRandomizerFunc(t *testing.T) {
	const name = "some-name"
	randomizedName := defaultNameRandomizerFunc(name)

	assert.Contains(t, randomizedName, name)
	assert.NotEqual(t, randomizedName, name)
}
