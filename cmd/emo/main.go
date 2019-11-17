package main

import (
	"fmt"
	"os"

	"github.com/syumai/emojidata"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "wanted emoji name is required. e.g. smile")
		return
	}
	emoji := emojidata.Get(os.Args[1])
	if emoji == nil {
		fmt.Fprintln(os.Stderr, "emoji was not found")
		return
	}
	fmt.Println(emoji.String())
}
