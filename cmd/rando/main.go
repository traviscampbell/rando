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
  i, insult     Generate a simple insult
  d, destroy    Generate an emotionally damaging insult
  c, codename   Generate a codename

Options for 'codename' only:
  -fword=""	Set the first word of the codename
  -lword=""	Set the last word of the codename
  -sep=""	Set custom seperator between the 2 words of the codename
  -imabitch	Ensure PC Prinpal would be on board with the generated name

Examples:
  rando insult "your mom"
  rando destroy rando
  rando codename -sep 💥
  rando destroy | cowsay -f $(ls /usr/share/cows | shuf -n1) | lolcat --spread 0.69
`

func main() {
	// made flag for codename cause it's the only one with actual options
	codenameCmd := flag.NewFlagSet("codename", flag.ExitOnError)

	fword := codenameCmd.String("fword", "", "set the first word of the codename")
	lword := codenameCmd.String("lword", "", "set the last word of the codename")
	sep := codenameCmd.String("sep", "", "custom seperator to use between the 2 words")
	// make user admit they're a bitch to gen a pro safe-spaces codename
	pussy := codenameCmd.Bool("imabitch", false, "ensure pc principal would be down with the generated name")

	// print help and exit if no command given
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(0)
	}

	switch os.Args[1] {
	case "i":
		fallthrough
	case "insult":
		name := ""
		if len(os.Args) >= 3 {
			name = os.Args[2]
		}
		fmt.Println(rando.Insult(name))

	case "d":
		fallthrough
	case "destroy":
		name := ""
		if len(os.Args) >= 3 {
			name = os.Args[2]
		}
		fmt.Println(rando.Destroy(name))

	case "c":
		fallthrough
	case "codename":
		codenameCmd.Parse(os.Args[2:])
		cn := rando.Codename()
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
		fmt.Printf("\nCongratz, You %s! I guess fucking up is kinda like your thing, huh?\n\n", rando.Codename().WithSeperator(" ").Please())
		os.Exit(1)
	}
}
