package main

import "strings"

const _defaultKey = "DCJ"

const _vowels = "AEIOUaeiou"

func encryptMessage(key string, message string) string {
	if key == "" {
		key = _defaultKey
	}

	if message == "" {
		return ""
	}

	var result strings.Builder
	for _, character := range message {
		if strings.ContainsRune(_vowels, character) {
			result.WriteString(key)
		}

		result.WriteRune(character)
	}

	return result.String()
}
