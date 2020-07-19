package main

import (
	"fmt"
	"log"
	"os"

	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
	"github.com/syumai/emojidata"
)

func main() {
	if len(os.Args) == 2 {
		emoji := emojidata.Get(os.Args[1])
		if emoji == nil {
			fmt.Fprintln(os.Stderr, "emoji was not found")
			return
		}
		fmt.Println(emoji.String())
		return
	}

	// fuzzyfind
	emojis := emojidata.EmojiData
	idx, err := fuzzyfinder.Find(
		emojis,
		func(i int) string {
			return emojis[i].ShortName
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return emojis[i].String()
		}))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(emojis[idx].String())
}
