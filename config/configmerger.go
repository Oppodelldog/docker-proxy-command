package config

func mergeCommand(baseCommand *CommandDefinition, overlayCommand *CommandDefinition) *CommandDefinition {
	mergedCommand := new(CommandDefinition)

	mergedCommand.Name = resolvePropertyString(baseCommand.Name, overlayCommand.Name)
	mergedCommand.EntryPoint = resolvePropertyString(baseCommand.EntryPoint, overlayCommand.EntryPoint)
	mergedCommand.Image = resolvePropertyString(baseCommand.Image, overlayCommand.Image)
	mergedCommand.WorkDir = resolvePropertyString(baseCommand.WorkDir, overlayCommand.WorkDir)
	mergedCommand.Network = resolvePropertyString(baseCommand.Network, overlayCommand.Network)

	mergedCommand.AddGroups = resolvePropertyBool(baseCommand.AddGroups, overlayCommand.AddGroups)
	mergedCommand.RemoveContainer = resolvePropertyBool(baseCommand.RemoveContainer, overlayCommand.RemoveContainer)
	mergedCommand.Impersonate = resolvePropertyBool(baseCommand.Impersonate, overlayCommand.Impersonate)
	mergedCommand.IsInteractive = resolvePropertyBool(baseCommand.IsInteractive, overlayCommand.IsInteractive)

	mergedCommand.Volumes = resolvePropertyStringArray(baseCommand.Volumes, overlayCommand.Volumes)
	mergedCommand.EnvVars = resolvePropertyStringArray(baseCommand.EnvVars, overlayCommand.EnvVars)
	mergedCommand.Ports = resolvePropertyStringArray(baseCommand.Ports, overlayCommand.Ports)
	mergedCommand.AdditionalArgs = resolvePropertyStringArray(baseCommand.AdditionalArgs, overlayCommand.AdditionalArgs)

	mergedCommand.ReplaceArgs = resolvePropertyStringArray2D(baseCommand.ReplaceArgs, overlayCommand.ReplaceArgs)

	return mergedCommand
}

func resolvePropertyBool(bBase *bool, bOverlay *bool) *bool {
	var b bool

	if bBase != nil {
		b = *bBase
	}

	if bOverlay != nil {
		b = *bOverlay
	}

	return &b
}

func resolvePropertyString(sBase *string, sOverlay *string) *string {
	var s string

	if sBase != nil {
		s = *sBase
	}

	if sOverlay != nil {
		s = *sOverlay
	}

	return &s
}

func resolvePropertyStringArray(sBase *[]string, sOverlay *[]string) *[]string {
	var s []string

	if sBase != nil {
		s = *sBase
	}

	if sOverlay != nil {
		s = *sOverlay
	}

	return &s
}

func resolvePropertyStringArray2D(sBase *[][]string, sOverlay *[][]string) *[][]string {
	var s [][]string

	if sBase != nil {
		s = *sBase
	}

	if sOverlay != nil {
		s = *sOverlay
	}

	return &s
}
