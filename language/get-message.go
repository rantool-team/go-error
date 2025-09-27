package language

import "fmt"

func (m MessageSet) GetMessage(languageName string) string {
	if !m.HasThisLanguage(languageName) {
		panicBecuseNotExistLanguage(languageName)
	}

	languageId, ok := getLanguageIotaAndReturnIfHasPredefinedLanguage(languageName)
	if ok {
		return m.getMessageOnId(languageId)
	}
	return m.getInAnotherLanguages(languageName)
}

func (m MessageSet) HasThisLanguage(languageName string) bool {
	if _, exists := getLanguageIotaAndReturnIfHasPredefinedLanguage(languageName); exists {
		return true
	}

	if m.OtherLanguages != nil {
		return false
	}

	_, ok := m.OtherLanguages[languageName]

	return ok
}

func panicBecuseNotExistLanguage(languageName string) {
	panic(fmt.Sprintf("Error: The language %q don't exists", languageName))
}

func (m MessageSet) getMessageOnId(id Language) string {
	switch id {
	case ENGLISH:
		return m.English
	case PORTUGUESE:
		return m.Portuguese
	case SPANISH:
		return m.Spanish
	case MANDARIN_CHINESE:
		return m.MandarinChinese
	case HINDI:
		return m.Hindi
	case FRENCH:
		return m.French
	case ARABIC:
		return m.Arabic
	case BENGALI:
		return m.Bengali
	case RUSSIAN:
		return m.Russian
	case URDU:
		return m.Urdu
	case INDONESIAN:
		return m.Indonesian
	case GERMAN:
		return m.German
	case JAPANESE:
		return m.Japanese
	case SWAHILI:
		return m.Swahili
	case PUNJABI:
		return m.Punjabi
	case MARATHI:
		return m.Marathi
	case TELUGU:
		return m.Telugu
	case TURKISH:
		return m.Turkish
	case KOREAN:
		return m.Korean
	case VIETNAMESE:
		return m.Vietnamese
	}

	return m.GetMessage(DefaultLanguage)
}

func (m MessageSet) getInAnotherLanguages(name string) string {
	if m.OtherLanguages == nil {
		return m.GetMessage(DefaultLanguage)
	}

	message, ok := m.OtherLanguages[name]

	if !ok {
		return m.GetMessage(DefaultLanguage)
	}

	return message
}

func getLanguageIotaAndReturnIfHasPredefinedLanguage(languageName string) (Language, bool) {
	l, ok := LanguageNameToCodeMap[languageName]

	return l, ok
}
