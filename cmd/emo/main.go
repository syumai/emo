package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"time"

	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
	"github.com/syumai/emo"
)

var (
	showRandom              = flag.Bool("rand", false, "show random selected emoji")
	listSubcategories       = flag.Bool("listsub", false, "list emoji subcategories")
	showRandomBySubcategory = flag.String("randsub", "", "show random selected emoji by specified subcategory")
	findBySubcategory       = flag.Bool("findsub", false, "fuzzy find by subcategory of emoji")
)

func main() {
	flag.Parse()
	rand.Seed(time.Now().Unix())
	if *showRandom {
		i := rand.Intn(len(emo.EmojiData))
		emoji := emo.EmojiData[i]
		fmt.Print(emoji.String())
		return
	}

	if *showRandomBySubcategory != "" {
		subCat := *showRandomBySubcategory
		var subCatEmojis []*emo.Emoji
		for _, emoji := range emo.EmojiData {
			if emoji.Subcategory == subCat {
				subCatEmojis = append(subCatEmojis, emoji)
			}
		}
		if len(subCatEmojis) == 0 {
			fmt.Fprintf(os.Stderr, "subcategory %q not found", subCat)
			return
		}
		i := rand.Intn(len(subCatEmojis))
		emoji := subCatEmojis[i]
		fmt.Print(emoji.String())
		return
	}

	if *listSubcategories {
		subCatsMap := make(map[string]struct{})
		for _, emoji := range emo.EmojiData {
			subCatsMap[emoji.Subcategory] = struct{}{}
		}
		var subCats []string
		for subCat := range subCatsMap {
			subCats = append(subCats, subCat)
		}
		sort.Strings(subCats)
		for _, subCat := range subCats {
			fmt.Println(subCat)
		}
		return
	}

	if len(flag.Args()) == 1 {
		emoji := emo.Get(flag.Args()[0])
		if emoji == nil {
			fmt.Fprintln(os.Stderr, "emoji was not found")
			return
		}
		fmt.Print(emoji.String())
		return
	}

	// fuzzyfind
	emojis := emo.EmojiData
	idx, err := fuzzyfinder.Find(
		emojis,
		func(i int) string {
			if *findBySubcategory {
				return fmt.Sprintf("%s_%s", emojis[i].Subcategory, emojis[i].ShortName)
			}
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
	fmt.Print(emojis[idx].String())
}
