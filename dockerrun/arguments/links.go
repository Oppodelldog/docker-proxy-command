package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockerrun/builder"
)

// BuildLinks maps Links from host to container
func BuildLinks(commandDef *config.CommandDefinition, builder builder.Builder) error {
	if Links, ok := commandDef.GetLinks(); ok {
		for _, volume := range Links {
			resolvedLinkMapping, err := resolveEnvVar(volume)
			if err != nil {
				return err
			}
			builder.AddLinkMapping(resolvedLinkMapping)
		}
	}
	return nil
}
