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

// you no fuck me; rando a fuck-ah you
var youTryFuckMe = regexp.MustCompile(`(?i)tcam|tra(vis|cam)`)

// Insult generates a simple insult of the form:
// (<Name> is|You are) a(n)? <Adv> <Adj> <Noun>.
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
	if regexp.MustCompile("^[aeiou]").MatchString(randyAdj) {
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

// CN represents options that can be used
// to configure a codename generator.
type CN struct {
	fword string
	lword string
	sep   string
	clean bool
}

// NewCodenamer returns factory builder instance
// used to genrate and configure codenames.
func NewCodenamer() *CN {
	return &CN{}
}

// WithFirstWord sets the firstword to the word provided.
func (cn *CN) WithFirstWord(fw string) *CN {
	cn.fword = fw
	return cn
}

// WithLastWord sets the lastword to the word provided.
func (cn *CN) WithLastWord(lw string) *CN {
	cn.lword = lw
	return cn
}

// WithSeperator lets you customize what, if anything, will seperate the parts
func (cn *CN) WithSeperator(sep string) *CN {
	cn.sep = sep
	return cn
}

// WhatAPussy will make sure the codename is as much a lil bitch as you are
func (cn *CN) WhatAPussy() *CN {
	cn.clean = true
	return cn
}

// Please just generate the fucking codename already.
func (cn *CN) Please() string {
	s1, s2 := genCodename(cn.clean)

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
func genCodename(isAPussy bool) (string, string) {
	// aids helps clean-up overzealous nouns..
	aids := regexp.MustCompile(`piece of |sack of |-)`)
	n := aids.ReplaceAllString(randy(noun), "")

	// clean adj %50 of the time, everytime.
	// unless, they are an admitted pussy
	a := randy(adj)
	if isAPussy || randomInt(99)%2 == 0 {
		a = randy(cleanAdj)
	}

	return a, n
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
