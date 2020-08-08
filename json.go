//go:generate statik -src=./assets
package emojidata

import (
	"io"
	"io/ioutil"

	"github.com/goccy/go-json"
	"github.com/rakyll/statik/fs"
	_ "github.com/syumai/emojidata/statik"
)

func unmarshalEmojiDataJSON() ([]*Emoji, error) {
	var emojis []*Emoji
	emojiDataFile, err := openDataFile()
	defer emojiDataFile.Close()
	if err != nil {
		return nil, err
	}
	blob, err := ioutil.ReadAll(emojiDataFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(blob, &emojis)
	if err != nil {
		return nil, err
	}
	return emojis, err
}

func openDataFile() (io.ReadCloser, error) {
	staticFS, err := fs.New()
	if err != nil {
		return nil, err
	}
	return staticFS.Open("/emoji-data/emoji_pretty.json")
}
