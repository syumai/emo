# emojidata

- emojidata is a Go package to provide emoji data of https://github.com/iamcal/emoji-data

## Usage

```go
// EmojiData is a list of emoji
emoji := emojidata.EmojiData[0]

// print data of emoji
// { Name, Unified, NonQualified, Docomo, Au, Softbank, Google...
fmt.Println(emoji)

// get specified emoji by short name
starEmoji := emojidata.Get("star")

// print emoji as string: ⭐
fmt.Println(starEmoji.String())
```

## Install a CLI tool to get Emoji

- A CLI tool is given as [emo](https://github.com/syumai/emojidata/blob/master/cmd/emo).

### Installation

```
go install github.com/syumai/emojidata/cmd/emo@latest
```

### Usage of emo

#### Simple find

- To get emoji, **just give a name of emoji** as argument.

```
# Get emoji of exclamation mark.
$ emo exclamation
❗
```

#### Fuzzy find

* emo provides fuzzy find feature using [ktr0731/go-fuzzyfinder](https://github.com/ktr0731/go-fuzzyfinder) .

```
# Open fuzzy find window
$ emo
> exclamation
  3/1643
> excla # Enter to get emoji
❗
```

#### Random select

```
$ emo -rand
😁
```

#### Copying emoji

- To copy emoji to your clipboard, please use commands like pbcopy (on Mac) or xsel (on Linux).

```
# Copy emoji of star to clipboard.
$ emo star | pbcopy (or `xsel -ib` on Linux)
```

## License

- MIT

## Author

- this package: [syumai](https://github.com/syumai)
- emoji-data: [iamcal](https://github.com/iamcal)
