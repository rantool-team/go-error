package language

func (m MessageSet) AppendOnAllLanguages(appendFn func(languageName string, value string)) {
	m.appendOnAllPredefinedLanguages(appendFn)
	m.appendInNotPredefinedLanguages(appendFn)
}

func (m MessageSet) appendOnAllPredefinedLanguages(appendFn func(languageName string, value string)) {
	appendFn(ENGLISH_STRING, m.English)
	appendFn(PORTUGUESE_STRING, m.Portuguese)
	appendFn(SPANISH_STRING, m.Spanish)
	appendFn(MANDARIN_CHINESE_STRING, m.MandarinChinese)
	appendFn(HINDI_STRING, m.Hindi)
	appendFn(FRENCH_STRING, m.French)
	appendFn(ARABIC_STRING, m.Arabic)
	appendFn(BENGALI_STRING, m.Bengali)
	appendFn(RUSSIAN_STRING, m.Russian)
	appendFn(URDU_STRING, m.Urdu)
	appendFn(INDONESIAN_STRING, m.Indonesian)
	appendFn(GERMAN_STRING, m.German)
	appendFn(JAPANESE_STRING, m.Japanese)
	appendFn(SWAHILI_STRING, m.Swahili)
	appendFn(PUNJABI_STRING, m.Punjabi)
	appendFn(MARATHI_STRING, m.Marathi)
	appendFn(TELUGU_STRING, m.Telugu)
	appendFn(TURKISH_STRING, m.Turkish)
	appendFn(KOREAN_STRING, m.Korean)
	appendFn(VIETNAMESE_STRING, m.Vietnamese)
}

func (m MessageSet) appendInNotPredefinedLanguages(appendFn func(languageName string, value string)) {
	if m.OtherLanguages == nil {
		return
	}

	for language, value := range m.OtherLanguages {
		appendFn(language, value)
	}
}
