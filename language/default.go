package language

var defaultLanguage string = ENGLISH_STRING

func SetDefaultLanguage(languageName string) {
	if _, exists := getLanguageIotaAndReturnIfHasPredefinedLanguage(languageName); !exists {
		return
	}

	defaultLanguage = languageName
}
