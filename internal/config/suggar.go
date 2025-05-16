package config

type SuggarString struct {
	value string
}

func (ss *SuggarString) getString() string {
	return ss.value
}

func (ss *SuggarString) lineBreak() *SuggarString {
	ss.value += "\n"
	return ss
}

func (ss *SuggarString) tab() *SuggarString {
	ss.value += "\t"
	return ss
}

func (ss *SuggarString) appendString(str string) *SuggarString {
	ss.value += str
	return ss
}

func (ss *SuggarString) addSuggar(suggar SuggarString) *SuggarString {
	ss.value += suggar.value
	return ss
}
