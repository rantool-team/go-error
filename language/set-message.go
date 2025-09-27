package language

func (m *MessageSet) SetMessage(languageName string, newMessage string) {
	languageId, ok := getLanguageIotaAndReturnIfHasPredefinedLanguage(languageName)
	if ok {
		m.setMessageOnId(languageId, newMessage)
		return
	}

	m.setInAnotherLanguage(languageName, newMessage)
}

func (m *MessageSet) setMessageOnId(id Language, newMessage string) {
	switch id {
	case ENGLISH:
		m.English = newMessage
		return
	case PORTUGUESE:
		m.Portuguese = newMessage
		return
	case SPANISH:
		m.Spanish = newMessage
		return
	case MANDARIN_CHINESE:
		m.MandarinChinese = newMessage
		return
	case HINDI:
		m.Hindi = newMessage
		return
	case FRENCH:
		m.French = newMessage
		return
	case ARABIC:
		m.Arabic = newMessage
		return
	case BENGALI:
		m.Bengali = newMessage
		return
	case RUSSIAN:
		m.Russian = newMessage
		return
	case URDU:
		m.Urdu = newMessage
		return
	case INDONESIAN:
		m.Indonesian = newMessage
		return
	case GERMAN:
		m.German = newMessage
		return
	case JAPANESE:
		m.Japanese = newMessage
		return
	case SWAHILI:
		m.Swahili = newMessage
		return
	case PUNJABI:
		m.Punjabi = newMessage
		return
	case MARATHI:
		m.Marathi = newMessage
		return
	case TELUGU:
		m.Telugu = newMessage
		return
	case TURKISH:
		m.Turkish = newMessage
		return
	case KOREAN:
		m.Korean = newMessage
		return
	case VIETNAMESE:
		m.Vietnamese = newMessage
		return
	}

	m.SetMessage(defaultLanguage, newMessage)
}

func (m *MessageSet) setInAnotherLanguage(languageName string, newMessage string) {
	m.initializeLanguageIfNotInitialized()

	m.OtherLanguages[languageName] = newMessage
}

func (m *MessageSet) initializeLanguageIfNotInitialized() {
	if m.OtherLanguages == nil {
		m.OtherLanguages = make(map[string]string)
	}
}
