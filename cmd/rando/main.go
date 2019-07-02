package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/traviscampbell/rando"
)

var usage = `
Usage: rando COMMAND [target|opts...]

  ██████╗  █████╗ ███╗   ██╗██████╗  ██████╗
  ██╔══██╗██╔══██╗████╗  ██║██╔══██╗██╔═══██╗
  ██████╔╝███████║██╔██╗ ██║██║  ██║██║   ██║
  ██╔══██╗██╔══██║██║╚██╗██║██║  ██║██║   ██║
  ██║  ██║██║  ██║██║ ╚████║██████╔╝╚██████╔╝
  ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝╚═════╝  ╚═════╝

    Random Insult and Codename Generator

Commands:
  i, insult    Generate a simple insult
  d, destroy   Generate an emotionally damaging insult
  c, codename  Generate a codename

Options for 'codename' only:
  -fword="" Set the first word of the codename
  -lword="" Set the last word of the codename
  -sep=""   Set custom seperator between the 2 words of the codename
  -imabitch Ensure PC Prinpal/HR would be on board with the generated name

Bonus Fuckery:
  -easteregg="" RegExp rando will defend if any shade gets thrown at it [default: (?i)tcam|tra(vis|cam)]
  -poach=false  Disable the protective easter-egg shield

Examples:
  rando insult "your mom"
  rando destroy rando
  rando codename -sep 💥
  rando destroy | cowsay -f $(ls /usr/share/cowsay/cows | shuf -n1) | lolcat --spread 0.69
`

func printUsage() {
	fmt.Println(usage)
}

func main() {
	// flagset and subcommand flag pointers for insult/destroy commands
	idfagset := flag.NewFlagSet("", flag.ExitOnError)
	easterEgg := idfagset.String("easteregg", "", "regexp rando will defend if any shade gets thrown its way")
	poachEgg := idfagset.Bool("poach", false, "disable the protective easter egg shield")
	idfagset.Usage = printUsage

	// flagset and subcommand flag pointers for codename command
	cnfagset := flag.NewFlagSet("codename", flag.ExitOnError)
	fword := cnfagset.String("fword", "", "set the first word of the codename")
	lword := cnfagset.String("lword", "", "set the last word of the codename")
	sep := cnfagset.String("sep", "", "custom seperator to use between the 2 words")
	// make user admit they're a bitch to gen a pro safe-spaces codename
	pussy := cnfagset.Bool("imabitch", false, "ensure pc principal would be down with the generated name")
	cnfagset.Usage = printUsage

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(0)
	}

	idfagset.Parse(os.Args[1:])
	if len(*easterEgg) >= 1 {
		rando.SetEasterEggRegexp(*easterEgg)
	}
	if *poachEgg {
		rando.PoachEasterEgg()
	}

	name := ""
	if len(os.Args) >= 3 {
		name = os.Args[2]
	}

	switch os.Args[1] {
	case "i", "insult":	
		fmt.Println(rando.Insult(name))

	case "d", "destroy":
		fmt.Println(rando.Destroy(name))

	case "c", "codename":
		cnfagset.Parse(os.Args[2:])

		cn := rando.NewCodenamer()
		if *fword != "" {
			cn.WithFirstWord(*fword)
		}
		if *lword != "" {
			cn.WithLastWord(*lword)
		}
		if *sep != "" {
			cn.WithSeperator(*sep)
		}
		if *pussy {
			cn.WhatAPussy()
		}

		fmt.Println(cn.Please())

	default:
		fmt.Printf("\nCongratz, you %s! You fucked up. Try again I suppose...\n\n", rando.NewCodenamer().WithSeperator(" ").Please())
		fmt.Println(usage)
		os.Exit(1)
	}
}
