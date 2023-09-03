package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/traviscampbell/rando"
)

var banner = fmt.Sprintf(`
  ██████╗  █████╗ ███╗   ██╗██████╗  ██████╗
  ██╔══██╗██╔══██╗████╗  ██║██╔══██╗██╔═══██╗
  ██████╔╝███████║██╔██╗ ██║██║  ██║██║   ██║
  ██╔══██╗██╔══██║██║╚██╗██║██║  ██║██║   ██║
  ██║  ██║██║  ██║██║ ╚████║██████╔╝╚██████╔╝
  ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝╚═════╝  ╚═════╝

    Random Insult and Codename Generator

      ver. %s
`, version)

var usage = fmt.Sprintf(`  %s
  Usage: rando COMMAND [target|opts...]

  Commands:
    i, insult    Generate a simple insult
    d, destroy   Generate an emotionally damaging insult
    c, codename  Generate a codename

  Options for 'codename' only:
    -fword="" Set the first word of the codename
    -lword="" Set the last word of the codename
    -sep=""   Set custom seperator between the 2 words of the codename
    -imabitch Ensure HR/PC-Prinpal would be onboard with the generated name
    -caps     Output in ALL CAPS
    -half     Keep adj clean (1st part) and leave the noun filthy. Has its moments...

  Bonus Fuckery:
    -easteregg="(?i)t(cam|ra)(vis|cam)?" Regex that will cause rando to clap-back at any would be haters
    -poach=false                         Disable protective easter-egg shield so you can hurt my feelings

  Examples:
    rando insult your mom
    rando d github
    rando destroy | cowsay --random | sinebow
    rando codename -sep
    for _ in {1..5}; do rando c; done
`, banner)

func printUsage() {
	fmt.Println(usage)
}

var version = "(ﾉ☉ヮ⚆)ﾉ ⌒*:･ﾟ✧"

func main() {
	hflag := flag.Bool("h", false, "your looking at it. hope it helps...")
	helpflag := flag.Bool("help", false, "your looking at it. hope it helps...")
	vflag := flag.Bool("v", false, "display version banner and exit")
	versionflag := flag.Bool("version", false, "display version banner and exit")

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
	caps := cnfagset.Bool("caps", false, "print the codename in all CAPS")
	half := cnfagset.Bool("half", false, "keep the adj clean and keep the noun filthy. it has its moments, on occasion")
	// make user admit they're a bitch to gen a safe-space codename
	pussy := cnfagset.Bool("imabitch", false, "ensure pc principal would be down with the generated name")
	cnfagset.Usage = printUsage
	flag.Parse()

	if *vflag || *versionflag {
		fmt.Println(banner)
		os.Exit(0)
	}

	if len(os.Args) < 2 || *hflag || *helpflag {
		printUsage()
		os.Exit(0)
	}

	if err := idfagset.Parse(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(*easterEgg) >= 1 {
		rando.SetEasterEggRegexp(*easterEgg)
	}
	if *poachEgg {
		rando.PoachEasterEgg()
	}

	name := ""
	if len(os.Args) >= 3 {
		name = strings.Join(os.Args[2:], " ")
	}

	switch os.Args[1] {
	case "i", "insult":
		fmt.Println(rando.Insult(name))

	case "d", "destroy":
		fmt.Println(rando.Destroy(name))

	case "c", "codename":
		if err := cnfagset.Parse(os.Args[2:]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		cn := rando.NewCodename()
		if *fword != "" {
			cn.WithFWord(*fword)
		}
		if *lword != "" {
			cn.WithLWord(*lword)
		}
		if *sep != "" {
			cn.WithSep(*sep)
		}
		if *pussy {
			cn.Pussify()
		}
		if *caps {
			cn.WithCAPS()
		}
		if *half {
			cn.HalfClean()
		}

		fmt.Println(cn.Generate())

	default:
		fmt.Printf("Congratz, ya done goofed, you %s!\nTry running it with one of the actual command values like: rando (i|d|c)\n", strings.ToLower(rando.NewCodename().WithSep(" ").Generate()))
		os.Exit(69)
	}
}
