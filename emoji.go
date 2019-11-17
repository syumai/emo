package emojidata

import (
	"fmt"
	"strconv"
	"strings"
)

type Emoji struct {
	Name            string   `json:"name"`
	Unified         string   `json:"unified"`
	NonQualified    string   `json:"non_qualified"`
	Docomo          string   `json:"docomo"`
	Au              string   `json:"au"`
	Softbank        string   `json:"softbank"`
	Google          string   `json:"google"`
	Image           string   `json:"image"`
	SheetX          int      `json:"sheet_x"`
	SheetY          int      `json:"sheet_y"`
	ShortName       string   `json:"short_name"`
	ShortNames      []string `json:"short_names"`
	Text            string   `json:"text"`
	Texts           []string `json:"texts"`
	Category        string   `json:"category"`
	SortOrder       int      `json:"sort_order"`
	AddedIn         string   `json:"added_in"`
	HasImgApple     bool     `json:"has_img_apple"`
	HasImgGoogle    bool     `json:"has_img_google"`
	HasImgTwitter   bool     `json:"has_img_twitter"`
	HasImgFacebook  bool     `json:"has_img_facebook"`
	HasImgMessenger bool     `json:"has_img_messenger"`
}

var EmojiData []*Emoji
var emojiMap map[string]*Emoji

func init() {
	emojis, err := unmarshalEmojiDataJSON()
	if err != nil {
		panic(err)
	}
	EmojiData = emojis
	emojiMap = make(map[string]*Emoji, len(emojis))
	for _, emoji := range emojis {
		if emoji.ShortName == "" {
			continue
		}
		emojiMap[emoji.ShortName] = emoji
	}
}

func (e *Emoji) String() string {
	parts := strings.Split(e.Unified, "-")
	for i, p := range parts {
		if len(p) > 4 {
			parts[i] = fmt.Sprintf("\\U%08s", p)
		} else {
			parts[i] = "\\u" + p
		}
	}
	parts = append([]string{`"`}, parts...)
	parts = append(parts, `"`)
	result, _ := strconv.Unquote(strings.Join(parts, ""))
	return result
}
