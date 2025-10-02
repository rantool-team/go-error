package errorstruct

func (e Error) getLocalDefinition() string {
	if e.HasLocalDefined {
		lineChar := LocalDemonstrationLineChar(e.Context.Local)
		res := lineChar.GetMessage(e.Context.Language) + "\n"
		modPackFile := LocalDemonstrationFileModPackWorkspace(e.Context.Local)
		res += modPackFile.GetMessage(e.Context.Language)

		return res
	}

	return ""
}
