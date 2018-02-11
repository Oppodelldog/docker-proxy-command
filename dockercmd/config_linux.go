package dockercmd

import "github.com/Oppodelldog/docker-proxy-command/config"

func addGroups(commandDef *config.CommandDefinition,builder *Builder){
	if addGroups, ok := commandDef.GetAddGroups(); ok {
		err = buildGroups(addGroups, builder)
		if err != nil {
			return nil, err
		}
	}
}
