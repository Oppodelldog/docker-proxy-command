package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockerrun/builder"
)

func addImpersonation(commandDef *config.CommandDefinition, builder builder.Builder) error {
	_ = commandDef
	_ = builder

	return nil
}
