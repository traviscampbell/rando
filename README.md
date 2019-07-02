# rando - Random Insult and Codename Generator

![usage](img/usage.png)

This project was initially started as a means to make our slackbot at work a bit more interesting. It's design was heavily influenced by the glorious [Better Off Ted](https://www.youtube.com/watch?v=Bh7Nz4bIwss) episode in which they develop an ingenious Insult Formula.

It was quickly realized that there were many more useful applications for `rando`, such as:

- Making users feel bad about themselves when mistyping their passwords
- Enhancing boring errors by humilating the function that caused the error
- Populating DB values like addedby/createdby with funny offensive codenames
- Piping generated insults into `espeak` or `say` when clients are nearby
- Helping new hires develop tougher skin, and self-esteem issues, etc...

It can be used as command-line tool or imported as a library in other projects. Some example usuage is provided down yonder...

---

## Install

`$> go get github.com/traviscampbell/rando/...`

---

## CLI Usage Examples

__Basic Insult:__ `rando i, insult [name]`

Just a simple, succinct, top-shelf insult that goes well with any occasion.

```shell
$> rando i
You are a filthy reprehensible douchebag.

$> rando insult rando
Rando is a dreadfully toxic cuck.

$> rando insult "your mom"
Your Mom is a horribly incompetent thundercunt.
```

__Emotionally Crippling Insult:__ `rando d, destory [name]`

This generates a powerful compound insult that is then topped-off with finisher that is sure to cut the target to their core, keeping them awake at night reflecting on just how accurate the statement truly was.

TODO: Scrape social media/financial accts of the target to deeply personalize the insult ensuring maximum effect.

```shell
$> rando d
You are a ridiculously reprehensible fucknugget and a dense cocksucking shithead whose birth certificate is an apology from the condom factory.

$> rando destroy rando
Rando is an absurdly dense shitstain, and a parasitic soul-destroying cesspool of sub-human filth.
```

__Codename Generation:__ `rando c, codename [opts]`

- `-fword`: specifically set the first word of the codename (thanks @jakewarren)
- `-lword`: specifically set the second word of the codename (thanks @jakewarren)
- `-sep`: set a custom seperator between the codename words
- `-imabitch`: just admit you are a little bitch and it'll give you a clean pc safe-space friendly name

```shell
$> rando codename
CretinousFucknugget

$> rando c -fword fappy
FappyScrotum

$> rando codename -sep 💥
Sinister💥Schlong

$> rando c -imabitch
LoudPanda
```

---

## Library Usage Example

```go
package main

import (
    "fmt"

    "github.com/traviscampbell/rando"
) 

func main() {
    // submiting an empty string to insult or destroy
    // will default the insult to the prefix "You are"
    fmt.Println(rando.Insult(""))
    fmt.Println(rando.Destoy(""))

    // may or may not destroy me ⊙▃⊙
    fmt.Println(rando.Destroy("travis"))

    // (ﾉ☉ヮ⚆)ﾉ ⌒*:･ﾟ✧ codenames
    cn := rando.Codenamer()
    fmt.Println(cn.Please())
    fmt.Println(cn.WithSeperator("𓂸").Please())
}
```

---

### Bonus

Keep your self-esteem in check by adding pretty self-affirmations to your `.bashrc` file:

```shell
$> rando destroy | cowsay -f $(ls /usr/share/cowsay/cows | shuf -n1) | lolcat --spread 0.69
```

![meanrainbowcow](img/meanRainbowCow.png)
