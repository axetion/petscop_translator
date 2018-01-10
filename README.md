# petscop_translator

This is a simple tool that "translates" the Petscop phonetic button code.

Prounciation data (`dictionary.gob`) from [The CMU Pronouncing Dictionary.](http://www.speech.cs.cmu.edu/cgi-bin/cmudict)
Word frequency data (`frequencies.gob`) from a sample of the [Corpus of Contemporary American English.](https://www.wordfrequency.info/)
Buttons -> Phonemes based on [/u/LittlestPetscop's post](https://www.reddit.com/r/Petscop/comments/7m27kf/the_typing_system_is_phonetic_heres_a_hopefully/)

## Building

No automated builds yet.

- Get [Go.](https://golang.org/)

- Clone the repository (if you don't have Git, you can download a ZIP using the button at the top.) If you clone from the command line, be sure to use `--recursive` or execute `git submodule init && git submodule update` after.

- Open up a terminal and `cd` to the resulting folder.

- If you're on Windows, run `set "GOPATH=%cd%"`. Otherwise (OS X, Linux) run  `` export GOPATH=`pwd` ``

- Run `go build petscop_translator.go`.

- Done. You can start it with `petscop_translator` (Windows) or `./petscop_translator`.

## Usage

Each individual button combination must be separated by spaces, and the buttons making up the combination can be separated by any form of punctuation. Surrounding parenthesis are optional.

Example:

```
$ ./petscop_translator
Loading dictionary...
---
Ask? (L1+X) (R1+TRIANGLE) (L2+X) (R2+SQUARE) (START) (L2+TRIANGLE)
> SIMPLE (exact match)

Ask? (R2+UP) (L2+TRIANGLE) (R1+TRIANGLE) (R2+SQUARE) (L1+X)
> FLIPS (exact match)

Ask? quit
```

## Why?

I made this for a few reasons, some more equal than the others:

1. To help assist in future translation efforts -- there's no way this code was a one-off given how much detail was obviously put in. A verbose mode `-v` is available to show the phonemes and all possibilities rather than just picking one match, which may be useful for this. 
2. To illustrate a few interesting points about the conversation in Petscop 11.
3. (most important) To demonstrate that it is in fact possible for a program to do this automatically. While I agree that it takes some suspension of belief for a PS1 game to be capable of this, given the memory consumption of the dictionary (about 5 MB uncompressed, which is more RAM than the PS1 has. Streaming from disk is possible, but would be very slow), a PC game can obviously do this. Not counting the stemmer which is an external library, this proof of concept is less than 500 lines of Go.

## How does it do?

Here's the output of the program running on the transcript of the Petscop 11 conversation:

```
Ask? (L2+DOWN) (LEFT) (L2+TRIANGLE) (R1+UP)
> HELLO (exact match)

Ask? (R2+SQUARE) (TRIANGLE) (L2+TRIANGLE)
> PALL (exact match)

Ask? (L2+DOWN) (LEFT) (L2+TRIANGLE)
> HELL (exact match)

Ask? (L2+DOWN) (LEFT) (L2+TRIANGLE) (R1+UP)
> HELLO (exact match)

Ask? (R2+UP) (L2+LEFT) (L2+SQUARE) (R1+CIRCLE)
> FUNNY (exact match)

Ask? (L2+DOWN) (X)
> HA (exact match)

Ask? (L2+DOWN) (X)
> HA (exact match)

Ask? (R2+SQUARE) (L2+TRIANGLE) (R1+X)
> PLAY (exact match)

Ask? (L2+X) (R1+CIRCLE) (R2+X) (L1+SQUARE) (R1+TRIANGLE) (L1+DOWN)
> MUSIC (approximate match)

Ask? (R2+UP) (TRIANGLE) (L2+CIRCLE)
> FOR (exact match)

Ask? (R2+TRIANGLE) (R1+X) (R2+TRIANGLE) (R1+CIRCLE)
> BABY (exact match)

Ask? (L1+TRIANGLE) (R1+CIRCLE)
> SHE (exact match)

Ask? (L2+START) (R1+TRIANGLE) (L2+TRIANGLE)
> WILL (exact match)

Ask? (R2+TRIANGLE) (R1+CIRCLE) (L1+DOWN) (START) (L2+X)
> BECOME (exact match)

Ask? (L2+X) (LEFT) (L2+TRIANGLE) (START) (R2+START) (R1+CIRCLE)
> MELODY (exact match)

Ask? (R2+LEFT) (SQUARE) (L2+SQUARE) (L1+DOWN) (L1+X)
> THANKS (approximate match)

Ask? (L2+TRIANGLE) (START) (R2+DOWN) (L2+TRIANGLE) (R1+CIRCLE)
> LOVELY (exact match)

Ask? (L2+TRIANGLE) (START) (R2+DOWN) (L2+TRIANGLE) (R1+CIRCLE)
> LOVELY (exact match)

Ask? (L1+TRIANGLE) (R1+CIRCLE)
> SHE (exact match)

Ask? (L1+START) (L2+CIRCLE) (R1+TRIANGLE) (R2+[SQUARE]) ([R2+CIRCLE])
> TRIPPED (approximate match)

Ask? (R1+X) (L2+SQUARE) (R2+[START])
> AND (approximate match)

Ask? (R2+UP) (LEFT) (L2+TRIANGLE)
> FELL (exact match)

Ask? (R1+X) (L2+SQUARE) (R2+[START])
> AND (approximate match)

Ask? (R1+TRIANGLE) (L1+X)
> IS (exact match)

Ask? (L2+TRIANGLE) (TRIANGLE) (L1+X) (R2+[CIRCLE])
> LOST (exact match)

Ask? (L1+X) (R2+CIRCLE) (X) (R2+SQUARE)
> STOP (exact match)

Ask? (L1+X) (X) (L2+CIRCLE) (R1+CIRCLE)
> SORRY (exact match)

Ask? (R2+CIRCLE) (R1+CIRCLE) (X) (L2+CIRCLE) (START)
> TIARA (exact match)

Ask? (R2+SQUARE) (L2+TRIANGLE) (R1+X) (L1+X)
> PLACE (exact match)

Ask? (R2+TRIANGLE) (SQUARE) (R2+START)
> BAD (exact match)

Ask? (L2+X) (R1+CIRCLE) (R2+X) (L1+SQUARE) (R1+TRIANGLE) (L1+DOWN)
> MUSIC (approximate match)

Ask? (R2+CIRCLE) (R2+X)
> TO (exact match)

Ask? (R2+START) (R2+X)
> DO (exact match)

Ask? (R1+TRIANGLE) (R2+CIRCLE)
> IT (exact match)

Ask? (L2+CIRCLE) (UP) (R2+CIRCLE)
> RIGHT (exact match)

Ask? (L2+SQUARE) (LEFT) (L1+DOWN) (L1+X) (R2+CIRCLE)
> NEXT (exact match)

Ask? (R2+CIRCLE) (UP) (L2+X)
> TIME (exact match)

Ask? (L1+X) (SQUARE) (R2+START)
> SAD (exact match)

Ask? (R1+UP) (L1+DOWN) (R1+X)
> OK (exact match)

Ask? (R2+SQUARE) (TRIANGLE) (L2+TRIANGLE)
> PALL (exact match)
```

Two things to note:

1. One of the first things that surprised me was that even thought PAUL is in the dictionary, even the earliest version of this program picked PALL just like Petscop itself. At least in this dictionary, PAUL and PALL are not homonyms as previously theorized: PAUL is `P AO L` (using the phonemes the CMU dictionary uses), whereas PALL is `P AA L`. I think it's pretty likely it was a delibrate "typo," but I also think it raises the possibility that if Petscop is generating these automatically, it could actually be using the CMU dictionary. I looked in my fairly old hardcover dictionary, for instance, and it had both PAUL and PALL as `P AO L`.

2. The "strange accent" /u/LittlestPetscop noted was a dick. The most noticable words that screwed it up were `AND` (expected `AH N D`, got `EY N D`), `BECOME` (expected `B IH K AH M`, got `B IY K AH M`, which is not *too* unusual of a pronounciation actually), `MUSIC` (expected `M Y UW Z IH K`, got `M IY UW Z IH K`), and `PLAYS` (expected `P L EY Z`, got `P L EY S`). The program employs fuzzy string matching, so it's able to tolerate 1 phoneme being off. When choosing between potential matches, it goes for the word with the most common stem (at least, of the frequency data I had available which was not much). This, along with allowing buttons like L1+X to be *either* an S or a Z, fixed most everything except for `PLAYS`. Unfortunately, `PLACE` is more common than `PLAY` in the data I used (by a tiny amount), and so it picks it. This is the only mistake it makes, and some future improvements could fix this.

## How does it work? (technical)

The dictionary is a BK-tree (https://en.wikipedia.org/wiki/BK-tree) mapping lists of the phonemes to words with that pronunciation. I used standard Levenshtein distance as the metric and the default threshold when searching is a distance of 1, meaning that you can add, delete, or change one of the phonemes and still have a match. This seems to work pretty well.

The frequency data is just a normal hash table matching *stems* (using Porter stemming) to their rank from the word frequency list I got. Both of these are constructed by `make_dict.go`, which then serializes them in delightful Go `.gob` files.

The main program `petscop_translator.go` reads in a line of button combinations and parses them (`query.go:ParseButtons`) into a list of 2-tuples like `{"L2", "X"}, {"START"}...`. (I used pretty much the same syntax /u/LittlestPetscop did, including calling the cross X because it's faster to type.)

It then maps the buttons into *possible* ARPAbet phonemes using the mapping in `arpabet.go`. I say *possible* because to account for the strange accent, a few of the buttons have multiple possible phonemes. When computing the Levenshtein distance, anything that equals one of the possibilities is a match (no substitution needed).

After that, it queries the BK-tree (`bktree.go`) and gets back a list of matches and their Levenshtein distances away from the input key. I pick the best match (`query.go:PickMatch`) by favoring lower Levenshtein distances, and tie breaking based on which word had a more frequent stem in the frequency data.

## Future improvements

- Better frequency data. Using just individual words is kind of stupid, n-grams would be much better. (considers previous words)
- Part of speech tagging. This could be used to better pick a single match, e.g. it's unlikely that a verb will come after another in a sentence.
- Travis & AppVeyor, so you don't have to build it by hand...
- A webapp version using GopherJS.
