package language

import "slices"

func (l *MessageSet) HasContentInSomeLanguage(content string) bool {
	languages := l.GetAllLanguages()

	return slices.Contains(languages, content)
}