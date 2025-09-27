package language

const (
	ENGLISH = iota
	PORTUGUESE
	SPANISH
	MANDARIN_CHINESE
	HINDI
	FRENCH
	ARABIC
	BENGALI
	RUSSIAN
	URDU
	INDONESIAN
	GERMAN
	JAPANESE
	SWAHILI
	PUNJABI
	MARATHI
	TELUGU
	TURKISH
	KOREAN
	VIETNAMESE

	ENGLISH_STRING          = "English"
	PORTUGUESE_STRING       = "Portuguese"
	SPANISH_STRING          = "Spanish"
	MANDARIN_CHINESE_STRING = "Mandarin Chinese"
	HINDI_STRING            = "Hindi"
	FRENCH_STRING           = "French"
	ARABIC_STRING           = "Arabic"
	BENGALI_STRING          = "Bengali"
	RUSSIAN_STRING          = "Russian"
	URDU_STRING             = "Urdu"
	INDONESIAN_STRING       = "Indonesian"
	GERMAN_STRING           = "German"
	JAPANESE_STRING         = "Japanese"
	SWAHILI_STRING          = "Swahili"
	PUNJABI_STRING          = "Punjabi"
	MARATHI_STRING          = "Marathi"
	TELUGU_STRING           = "Telugu"
	TURKISH_STRING          = "Turkish"
	KOREAN_STRING           = "Korean"
	VIETNAMESE_STRING       = "Vietnamese"
)

type Language int

var LanguageNameToCodeMap = map[string]Language{
	ENGLISH_STRING:          ENGLISH,
	PORTUGUESE_STRING:       PORTUGUESE,
	SPANISH_STRING:          SPANISH,
	MANDARIN_CHINESE_STRING: MANDARIN_CHINESE,
	HINDI_STRING:            HINDI,
	FRENCH_STRING:           FRENCH,
	ARABIC_STRING:           ARABIC,
	BENGALI_STRING:          BENGALI,
	RUSSIAN_STRING:          RUSSIAN,
	URDU_STRING:             URDU,
	INDONESIAN_STRING:       INDONESIAN,
	GERMAN_STRING:           GERMAN,
	JAPANESE_STRING:         JAPANESE,
	SWAHILI_STRING:          SWAHILI,
	PUNJABI_STRING:          PUNJABI,
	MARATHI_STRING:          MARATHI,
	TELUGU_STRING:           TELUGU,
	TURKISH_STRING:          TURKISH,
	KOREAN_STRING:           KOREAN,
	VIETNAMESE_STRING:       VIETNAMESE,
}
