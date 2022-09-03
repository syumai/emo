package emo

import (
	"strconv"
	"strings"
)

func Get(key string) *Emoji {
	emoji, ok := emojiMap[key]
	if !ok {
		return nil
	}
	return emoji
}

// GetFromEmojiString gets Emoji data by single emoji character.
// If no Emoji data is found, returns nil.
func GetFromEmojiString(emoji string) *Emoji {
	runes := []rune(emoji)
	codes := make([]string, len(runes))
	for i, r := range runes {
		codes[i] = strings.ToUpper(strconv.FormatInt(int64(r), 16))
	}
	unified := strings.Join(codes, "-")
	return unifiedMap[unified]
}
