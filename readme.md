# Cardtage
An attempt at using the swiss army knife MagickWand libraries to make printing playing cards easier.

# Usage
```bash
cardtage -h
Cardtage does one thing and it does that one thing really-mediocerly.
                                Give it files, tell it how far to space them, and tell it the size of page
                                you'd like to make (like 8.5x11) and tada! A montage of those images.
                                Useful for printing diy playing cards

Usage:
  cardtage [flags]
  cardtage [command]

Available Commands:
  help        Help about any command
  version     Print the version number of cardtage

Flags:
  -b, --border float    in inches, border size around each card (default 0.01)
  -c, --card string     in inches, the size of the card/tile (default "2.5x3.5")
  -d, --density float   PPI/resolution to use when ingesting and exporting (default 100)
  -h, --help            help for cardtage
  -m, --margin string   in inches, the margin that should be respected (default "0.25x0.25")
  -p, --page string     in inches, size of the page you'd like to output (default "8.5x11")

Use "cardtage [command] --help" for more information about a command.
```

# Examples
1. Download
2. Convert 
```bash
$ cardtage "input/*" -c "2.64x1.73" -b 0.01 -p "8.5x11" -m "0.25x0.25" -d 200 out.pdf
Writing montage to: 'out.pdf'...
Complete.
```


