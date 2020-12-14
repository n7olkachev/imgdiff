# imgdiff

Faster than [the fastest in the world pixel-by-pixel image difference tool](https://github.com/dmtrKovalenko/odiff).

## Why?

imgdiff isn't as fast as a tool like this should be and I'm not proud of it, but it is 3X faster than
[the fastest in the world pixel-by-pixel image difference tool](https://github.com/dmtrKovalenko/odiff),
so maybe you'll find it useful.

## Features

It can do everything [odiff](https://github.com/dmtrKovalenko/odiff) can. Faster.

## Benchmarks

I've tested it on Linux, Intel(R) Core(TM) i7-4700HQ CPU @ 2.40GHz, 8 cores.

[Cypress image](https://github.com/dmtrKovalenko/odiff/blob/main/images/www.cypress.io.png) 3446 x 10728

| Command                                                        |      Mean [s] | Min [s] | Max [s] | Relative |
| :------------------------------------------------------------- | ------------: | ------: | ------: | -------: |
| `imgdiff images/cypress-1.png images/cypress-2.png output.png` | 1.442 ± 0.012 |   1.420 |   1.462 |     1.00 |
| `odiff images/cypress-1.png images/cypress-2.png output.png`   | 6.475 ± 0.092 |   6.300 |   6.583 |     4.49 |

[Water image](https://github.com/dmtrKovalenko/odiff/blob/main/images/water-4k.png) 8400 x 4725

| Command                                                    |       Mean [s] | Min [s] | Max [s] | Relative |
| :--------------------------------------------------------- | -------------: | ------: | ------: | -------: |
| `imgdiff images/water-1.png images/water-2.png output.png` | 1.908 ± 0.0058 |   1.841 |   2.002 |     1.00 |
| `odiff images/water-1.png images/water-2.png output.png`   |  6.016 ± 0.415 |   5.643 |   7.140 |     3.15 |

## Usage

```
Usage: imgdiff [--threshold THRESHOLD] [--diff-image] [--fail-on-layout] BASE COMPARE OUTPUT

Positional arguments:
  BASE                   Base image.
  COMPARE                Image to compare with.
  OUTPUT                 Output image path.

Options:
  --threshold THRESHOLD, -t THRESHOLD
                         Color difference threshold (from 0 to 1). Less more precise. [default: 0.1]
  --diff-image           Render image to the diff output instead of transparent background. [default: false]
  --fail-on-layout       Do not compare images and produce output if images layout is different. [default: false]
  --help, -h             display this help and exit
```

## Download

You can find pre-built binaries [here](https://github.com/n7olkachev/imgdiff/releases/tag/v1.0.0).
imgdiff is written in Go, so there shouldn't be any troubles to compile it for the most of popular platforms.
