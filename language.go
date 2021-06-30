package ietf_language_tag

// GetISO6391 ISO 639-1
func GetISO6391(lang string) string { return iso6391Tag[lang] }

// GetISO31661 ISO 3166-1
func GetISO31661(lang string) string { return iso31661Region[lang] }
