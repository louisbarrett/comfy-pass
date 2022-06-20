# comfy-pass

A simple command-line tool for generating random words using characters from each hemisphere the keyboard.

## Usage

```
Usage of comfy-pass
  -l    left hand words
  -r    right hand words
  -n int
        Number of words to generate in pw (default 3)

```

This will generate 3 random words using letters from the left side of the keyboard.

```
go run main.go -n 3 -l
raxed-steward-cadge
```

This will generate 3 random words using letters from the right side of the keyboard.
```
go run main.go -n 3 -r
yum-oh-plum
```
## Options

`-n`: The number of words to generate. Defaults to 3.

`-l`: generate words for the left side of the keyboard.

`-r`: generate words for the right side of the keyboard.

