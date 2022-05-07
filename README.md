# emo

[![Go Reference](https://pkg.go.dev/badge/github.com/syumai/emo.svg)](https://pkg.go.dev/github.com/syumai/emo)

- emo is a CLI tool to fuzzy-find emojis.
  - This package is using provide emoji data of https://github.com/iamcal/emoji-data

## Installation

- This tool requires Go 1.18+ for installation.

```
go install github.com/syumai/emo/cmd/emo@latest
```

## Usage

### Simple find

```
$ emo exclamation
❗
```

### Fuzzy find

```
# Open fuzzy find window
$ emo
> exclamation
  3/1643
> excla # Enter to get emoji
❗
```

### Random select

```
$ emo -rand
😁
```

### Fuzzy find / Random select by subcategory

```
$ emo -listsub
alphanum
animal-amphibian
animal-bird
...
warning
writing
zodiac

$ emo -findsub
> warning_
  16/1817
> warning_no_bicycles # Enter to get emoji
🚳

$ emo -randsub animal-bird
🦜
```

### Copying emoji

- To copy emoji to your clipboard, please use commands like pbcopy (on Mac) or xsel (on Linux).

```
# Copy emoji of star to clipboard.
$ emo star | pbcopy (or `xsel -ib` on Linux)
```

## Using as library

### Installation

```console
go get github.com/syumai/emo
```

### Usage

```go
// EmojiData is a list of emoji
emoji := emo.EmojiData[0]

// print data of emoji
// { Name, Unified, NonQualified, Docomo, Au, Softbank, Google...
fmt.Println(emoji)

// get specified emoji by short name
starEmoji := emo.Get("star")

// print emoji as string: ⭐
fmt.Println(starEmoji.String())
```

## License

- MIT

## Author

- this package: [syumai](https://github.com/syumai)
- emoji-data: [iamcal](https://github.com/iamcal)
