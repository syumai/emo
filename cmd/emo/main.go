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
	emoj := emojidata.Get(os.Args[1])
	if emoj == nil {
		fmt.Fprintln(os.Stderr, "emoji was not found")
		return
	}
	fmt.Println(emoj.String())
}
