package emojidata

func Get(key string) *Emoji {
	emoji, ok := emojiMap[key]
	if !ok {
		return nil
	}
	return emoji
}
