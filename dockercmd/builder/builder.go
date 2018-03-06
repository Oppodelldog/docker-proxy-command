package builder

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

type (
	// builder can be used to build a docker run command
	builder struct {
		command         string
		subCommand      string
		imageName       string
		entryPoint      string
		network         []string
		args            []string
		portMappings    []string
		volumeMappings  []string
		envVarMappings  []string
		attachedStreams []string
		workingDir      []string
		containerName   []string
		addedGroups     []string
		containerUser   []string
		cmdArgs         []string
		stdIn           io.Reader
		stdOut          io.Writer
		stdErr          io.Writer

		buildArgs []string
	}
)

// NewBuilder returns a new docker command builder
func New() Builder {
	return &builder{
		command:    "docker",
		subCommand: "run",
		stdIn:      os.Stdin,
		stdOut:     os.Stdout,
		stdErr:     os.Stderr,
	}
}

// SetStdIn will be applied to the exec.Cmd
func (b *builder) SetStdIn(r io.Reader) Builder {
	b.stdIn = r

	return b
}

// SetStdOut will be applied to the exec.Cmd
func (b *builder) SetStdOut(w io.Writer) Builder {
	b.stdOut = w

	return b
}

// SetStdErr will be applied to the exec.Cmd
func (b *builder) SetStdErr(w io.Writer) Builder {
	b.stdErr = w

	return b
}

// AddPortMapping adds a mapping of ports between the docker container and the host
func (b *builder) AddPortMapping(portMapping string) Builder {
	b.portMappings = append(b.portMappings, "-p", portMapping)
	return b
}

// AddCmdArguments adds command arguments that are applied to the command executed inside the container
func (b *builder) AddCmdArguments(arguments []string) Builder {
	b.cmdArgs = append(b.cmdArgs, arguments...)
	return b
}

// AddArgument adds arguments to the docker run command
func (b *builder) AddArgument(argument string) Builder {
	b.args = append(b.args, argument)
	return b
}

// AttachTo attaches Streams to the docker-container.
// possible values: STDERR, STDOUT, STDIN
func (b *builder) AttachTo(stream string) Builder {
	b.attachedStreams = append(b.attachedStreams, "-a", stream)
	return b
}

// AddVolumeMapping adds a volulme mapping between the docker container and the host
func (b *builder) AddVolumeMapping(volume string) Builder {

	b.volumeMappings = append(b.volumeMappings, "-v", volume)
	return b
}

// AddEnvVar adds an environment variable to the docker-container.
// example: HOME=/home/myuser
func (b *builder) AddEnvVar(envVarDeclaration string) Builder {
	b.envVarMappings = append(b.envVarMappings, "-e", envVarDeclaration)
	return b
}

// AddGroup adds the given group name into the docker container.
func (b *builder) AddGroup(groupName string) Builder {
	b.addedGroups = append(b.addedGroups, "--group-add", groupName)
	return b
}

// SetEntryPoint sets the entry point for the docker run command
func (b *builder) SetEntryPoint(entryPoint string) Builder {
	b.entryPoint = entryPoint
	return b
}

// SetNetwork connects the docker container to the given docker-network
func (b *builder) SetNetwork(network string) Builder {
	b.network = []string{"--network", network}
	return b
}

// SetImageName sets the image on which base the container is created
func (b *builder) SetImageName(imageName string) Builder {
	b.imageName = imageName
	return b
}

// SetWorkingDir sets the default working dir for commands executed inside the container
func (b *builder) SetWorkingDir(workingDir string) Builder {
	b.workingDir = []string{"-w", workingDir}
	return b
}

// SetContainerName sets the display name of the container
func (b *builder) SetContainerName(containerName string) Builder {
	b.containerName = []string{"--name", containerName}

	return b
}

// SetContainerUserAndGroup sets the given userId:groupId as current user and group in the container
func (b *builder) SetContainerUserAndGroup(userID string, groupID string) Builder {
	b.containerUser = []string{"-u", fmt.Sprintf("%s:%s", userID, groupID)}

	return b
}

// Build builds the exec.Cmd which will start a docker-container
func (b *builder) Build() *exec.Cmd {

	cmd := exec.Command(b.command, b.subCommand)

	b.buildArgsAppend(b.args...)
	b.buildArgsAppend(b.containerName...)
	b.buildArgsAppend(b.workingDir...)
	b.buildArgsAppend(b.portMappings...)
	b.buildArgsAppend(b.volumeMappings...)
	b.buildArgsAppend(b.envVarMappings...)
	b.buildArgsAppend(b.addedGroups...)
	b.buildArgsAppend(b.containerUser...)
	b.buildArgsAppend(b.attachedStreams...)
	b.buildArgsAppend(b.network...)

	b.buildArgAppend(b.imageName)
	b.buildArgAppend(b.entryPoint)

	b.buildArgsAppend(b.cmdArgs...)

	if len(b.buildArgs) > 0 {
		cmd.Args = append(cmd.Args, b.buildArgs...)
	}

	cmd.Stdout = b.stdOut
	cmd.Stderr = b.stdErr
	cmd.Stdin = b.stdIn

	return cmd
}

func (b *builder) buildArgAppend(arg string) {
	b.buildArgs = append(b.buildArgs, arg)
}

func (b *builder) buildArgsAppend(args ...string) {
	b.buildArgs = append(b.buildArgs, args...)
}
