package language

func (m *MessageSet) SetMessage(languageName string, newMessage string) {
	languageId, ok := getLanguageIotaAndReturnIfHasPredefinedLanguage(languageName)
	if ok {
		m.setMessageOnId(languageId, newMessage)
	}

	m.setInAnotherLanguage(languageName, newMessage)
}

func (m *MessageSet) setMessageOnId(id Language, newMessage string) {
	switch id {
	case ENGLISH:
		m.English = newMessage
	case PORTUGUESE:
		m.Portuguese = newMessage
	case SPANISH:
		m.Spanish = newMessage
	case MANDARIN_CHINESE:
		m.MandarinChinese = newMessage
	case HINDI:
		m.Hindi = newMessage
	case FRENCH:
		m.French = newMessage
	case ARABIC:
		m.Arabic = newMessage
	case BENGALI:
		m.Bengali = newMessage
	case RUSSIAN:
		m.Russian = newMessage
	case URDU:
		m.Urdu = newMessage
	case INDONESIAN:
		m.Indonesian = newMessage
	case GERMAN:
		m.German = newMessage
	case JAPANESE:
		m.Japanese = newMessage
	case SWAHILI:
		m.Swahili = newMessage
	case PUNJABI:
		m.Punjabi = newMessage
	case MARATHI:
		m.Marathi = newMessage
	case TELUGU:
		m.Telugu = newMessage
	case TURKISH:
		m.Turkish = newMessage
	case KOREAN:
		m.Korean = newMessage
	case VIETNAMESE:
		m.Vietnamese = newMessage
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
