package rando

import (
	"crypto/rand" // don't act like you're not impressed...
	"fmt"
	"math/big"
	"regexp"
	"strings"
)

// randy string from given slice
func randy(source []string) string {
	return source[randomInt(len(source))]
}

// ensures name is properly formatted and/or defined
func formatName(name string) string {
	if name == "" {
		return "You are"
	}
	return strings.Title(name) + " is"
}

// Suprise, Motherfucker! Sunrise, Motherfucker! Some Fries, Motherfucker!
var youTryFuckMe = regexp.MustCompile(`(?i)tcam|tra(vis|cam)`)

// Insult generates a simple insult of the form:
// <Name> is|You are <Adv> <Adj> <Noun>.
func Insult(n string) string {
	name := formatName(n)
	if youTryFuckMe.MatchString(name) {
		return youNoFuckMe
	}
	return name + genInsult() + "."
}

func genInsult() string {
	return fmt.Sprintf(
		" %s %s %s",
		randy(adv),
		randy(adj),
		randy(noun),
	)
}

// Destroy them on an emotional level with a
// devestating compound insult in the form:
// <Insult>, and a(n) <Adj> <Desc> <Noun> <Finisher>.
func Destroy(n string) string {
	name := formatName(n)
	if youTryFuckMe.MatchString(name) {
		return youNoFuckMe
	}
	return name + genInsult() + genDestroyer() + "."
}

func genDestroyer() string {
	randyAdj := randy(adj)

	// fucking nazi's
	nazi := "a"
	if regexp.MustCompile("^(a|e|i|o|u)").MatchString(randyAdj) {
		nazi = "an"
	}

	return fmt.Sprintf(
		" and %s %s %s %s",
		nazi,
		randyAdj,
		randy(desc),
		// inject a randy noun for finisher's w/o one built-in
		strings.Replace(randy(finisher), "<noun>", randy(noun), -1),
	)
}

// CN represents a codename instance
type CN struct {
	fword string
	lword string

	sep string

	clean bool
}

// Codename represents an Adj + Noun combo
// like `SluttyBallsack`, for instance.
func Codename() *CN {
	return &CN{}
}

// Set specific word to be used as the first word
func (cn *CN) WithFirstWord(fw string) *CN {
	cn.fword = fw
	return cn
}

// Set specific word to be used as the last word
func (cn *CN) WithLastWord(lw string) *CN {
	cn.lword = lw
	return cn
}

// If you don't want the words smushed together
// this could probably help you.
func (cn *CN) WithSeperator(sep string) *CN {
	cn.sep = sep
	return cn
}

// Make sure the codename is as much of a lil pussy bitch as you are
func (cn *CN) WhatAPussy() *CN {
	cn.clean = true
	return cn
}

// Codename will actually be generated now that you've asked nicely.
func (cn *CN) Please() string {
	var s1, s2 string

	switch {
	case cn.clean:
		s1, s2 = getCleanName()
	default:
		s1, s2 = getCodename()
	}

	if cn.fword != "" {
		s1 = cn.fword
	}

	if cn.lword != "" {
		s2 = cn.lword
	}

	s1 = strings.Title(s1)
	s2 = strings.Title(s2)

	return s1 + cn.sep + s2
}

// returns a random codename combo. The noun is generally
// the offensive part here and can be enhanced by any adj.
// Thus, the adj's cleanliness shall be determined at random.
func getCodename() (string, string) {
	// aids helps clean-up overzealous nouns..
	aids := regexp.MustCompile(`(piece of |sack of )`)

	// (V) (°,,,,°) (V) - 50/50? ...Why not zoidberg!? (whoop. whoop. whooop.)
	if randomInt(99)%2 == 0 {
		return randy(adj), aids.ReplaceAllString(randy(noun), "")
	}
	return randy(cleanAdj), aids.ReplaceAllString(randy(noun), "")
}

// returns a random clean adj/noun combo
func getCleanName() (string, string) {
	return randy(cleanAdj), randy(cleanNoun)
}

// -------------

// get a crypto random int between 0 and n
func randomInt(n int) uint64 {
	switch {
	case n == 0:
		return uint64(0)
	case n < 0:
		n *= -1
	}

	num, _ := rand.Int(rand.Reader, big.NewInt(int64(n)))

	return num.Uint64()
}
