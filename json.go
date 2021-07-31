package emojidata

import (
	_ "embed"

	"github.com/goccy/go-json"
)

//go:embed assets/emoji-data/emoji_pretty.json
var emojiDataJSON []byte

func unmarshalEmojiDataJSON() ([]*Emoji, error) {
	var emojis []*Emoji
	err := json.Unmarshal(emojiDataJSON, &emojis)
	if err != nil {
		return nil, err
	}
	return emojis, err
}
