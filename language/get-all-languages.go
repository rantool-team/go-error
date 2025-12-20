package language

func (m *MessageSet) GetAllLanguages() []string {
	res := []string{}

	if m.Arabic != "" {
		res = append(res, m.Arabic)
	}

	if m.Bengali != "" {
		res = append(res, m.Bengali)
	}

	if m.English != "" {
		res = append(res, m.English)
	}

	if m.French != "" {
		res = append(res, m.French)
	}

	if m.German != "" {
		res = append(res, m.German)
	}

	if m.Hindi != "" {
		res = append(res, m.Hindi)
	}

	if m.Indonesian != "" {
		res = append(res, m.Indonesian)
	}

	if m.Japanese != "" {
		res = append(res, m.Japanese)
	}

	if m.Korean != "" {
		res = append(res, m.Korean)
	}

	if m.MandarinChinese != "" {
		res = append(res, m.MandarinChinese)
	}

	if m.Marathi != "" {
		res = append(res, m.Marathi)
	}

	if m.Portuguese != "" {
		res = append(res, m.Portuguese)
	}

	if m.Punjabi != "" {
		res = append(res, m.Punjabi)
	}

	if m.Russian != "" {
		res = append(res, m.Russian)
	}

	if m.Spanish != "" {
		res = append(res, m.Spanish)
	}

	if m.Swahili != "" {
		res = append(res, m.Swahili)
	}

	if m.Telugu != "" {
		res = append(res, m.Telugu)
	}

	if m.Turkish != "" {
		res = append(res, m.Turkish)
	}

	if m.Urdu != "" {
		res = append(res, m.Urdu)
	}

	if m.Vietnamese != "" {
		res = append(res, m.Vietnamese)
	}

	return res
}