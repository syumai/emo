package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
	"github.com/syumai/emojidata"
)

var showRandom = flag.Bool("rand", false, "show random selected emoji")

func main() {
	flag.Parse()
	if *showRandom {
		rand.Seed(time.Now().Unix())
		i := rand.Intn(len(emojidata.EmojiData))
		emoji := emojidata.EmojiData[i]
		fmt.Println(emoji.String())
		return
	}

	if len(flag.Args()) == 1 {
		emoji := emojidata.Get(flag.Args()[0])
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
