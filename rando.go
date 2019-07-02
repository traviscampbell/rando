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

// randyWithArticle returns a random string from a slice
// along with the indefinite article 'a' or 'an' prefixed
func randyWithArticle(source []string) string {
	str := randy(source)
	if regexp.MustCompile(`^[aeiou]`).MatchString(str) {
		return "an " + str
	}
	return "a " + str
}

// normaleyes normalizes name to title-cased unless
// empty which defaults to 'You are' cause they are.
func normaleyes(name string) string {
	if name == "" {
		return "You are"
	}
	return strings.Title(name) + " is"
}

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

// ----------------------------------------------------------------

// easterEgg is a protective eggshell that when matched will
// cause rando to stick-up for you and shit all over any egg heads
// who thought they'd be cute and try to egg you on.
type easterEgg struct {
	nameRE *regexp.Regexp
	enabled bool
}

func (e *easterEgg) protectMe(n string) bool {
	return e.nameRE.MatchString(n)
}

// SetEasterEggRegexp lets you override the regexp val to help
// protect yourself from getting beaten, cracked, or fried
func SetEasterEggRegexp(reStr string) {
	humpty.nameRE = regexp.MustCompile(reStr)
}

// PoachEasterEgg disables protective easter egg shell
func PoachEasterEgg() {
	humpty.enabled = false
}

// humpty is the default egg singleton preactivated to defend me because I'm
// a little chicken who can dish it out, but(t)? fuck taking it... ⊙▂⊙
var humpty = easterEgg{
	nameRE: regexp.MustCompile(`(?i)tcam|tra(vis|cam)`),
	enabled: true,
}

// answers the age old question: you try fuck me!? by comparing
// the set regexp value against the provided name. If it matches
// then fuckery is afoot and he'll hop down from the wall and start
// cracking some fucking shells.
func youTryFuckMe(n string) bool {
	return humpty.enabled && humpty.protectMe(n)
}

// P.S. - Why the fuck were the king's horses trying to put an egg back together?
//
//  > Humpty Dumpty sat on a wall,
//  > Humpty Dumpty had a great fall.
//  > All the king's horses and all the king's men
//  > Couldn't put Humpty together again.
//
// After reviewing the historical eggsploits, and egg-witness testimonies it is believed that the
// king's horses conspired together and laid out a plan to scramble and fry the alleged egg-white
// nationalist, known as Humpty-Dumpty.
// By all eggounts, Humpty was rotten to the yolk. He was known to have strong-ties to a poultry
// number of fowl hard-boiled eggtremists and was dyed up-and-down with confederate fleeg colorings.
// Though, it was Humpty's terrifrying charisma that allowed him to farm a sizeable carton of followers
// out by the wall where he spent countless hours shelling out his far-white propaganda, fertilizing his
// brethren with his deep-fryed indoctrination, and eggciting violence.
// Fortunately, the king's horses were already 'yolked-as-fuck', and had been since long before it was cool.
// It was this runny-yolkness that provided the clarity to see that Humpty had benedict for too long and
// eggtreme measures were going to be needed to stop his spread of hatred and salmonella denials if their
// local community was ever going to become truly free range.
// The decision was made that Humpty had to be poached soon before all the townyolks rage came to a boil,
// and so a plan was hatched. They knew it was going to be no over-easy task to make it appear as though
// it was just another blameless eggnog-induced eggcident. Nonetheless, they decided the reward vastly
// outweighed the whisk and eggecuted their plan. Against all odds, it was an overwhelming success and
// Humpty's previously unchicked reign of terror plummeted with him as he flailed about like a headless
// chicken until finally making a truly devasting shell-shattering impact with the cold, hard, and unforgiving
// ground.
// Rumor has it, that even the beaten Humpty himself believed it to have been an eggcident. This stems from
// an allegged eggchange between one of the king's horses and Humpty as the king's horses and his
// men were trying to put Humpty back together eggain. Unfortunately, there is no eggreed upon consesus as
// to the eggact words that ended up inyolking Humpty's response which were to become his infamous last words,
// "Omelleta, that one slide!" Humpty conceded in an seemingly ashamed manner."My balance is such a
// yoke. Though, it makes sense my lack of balance led to my downfall. And on the sunny-side up, it sure did
// crack me up and that just pickles me to death," he stated in a most uncharacteristically positive and
// somewhat ironic way as he panned his gaze across the range one final time.
//
// tl;dr - TIL, Humpty Dumpty was a nazi-sympathizer and tells the story of a hate-crime that was born out
// of love, shrouded in secrecy, and perpetuated by a group of yoke-ass horses who choose to unite and take
// a stand against the devilled egg himself whisking everything just to one day be free range.
//
// P.S.S.
// Spent longer on this eggpun-riddled fact-infused historically accurate conspiracy doctrine
// than the actual code... If nothing else, this can serve as an illustrative example of how
// to be productively unproductive which is yet another gem to add to my evergrowing list of
// of skills/strengths you generally don't want to put on your resume... (ﾉ☉ヮ⚆)ﾉ ⌒*:･ﾟ✧

// ----------------------------------------------------------------

// Insult generates a simple insult of the form:
// (<Name> is|You are) a(n)? <Adv> <Adj> <Noun>.
func Insult(name string) string {
	if youTryFuckMe(name) {
		return youNoFuckMe
	}
	return normaleyes(name) + genInsult() + "."
}

func genInsult() string {
	return fmt.Sprintf(
		" %s %s %s",
		randyWithArticle(adv),
		randy(adj),
		randy(noun),
	)
}

// Destroy them on an emotional level with a
// devestating compound insult in the form:
// <Insult>, and a(n) <Adj> <CompoundAdj> <Noun> <Finisher>.
func Destroy(name string) string {
	if youTryFuckMe(name) {
		return youNoFuckMe
	}

	name = normaleyes(name)
	return name + genDestroyer(name) + "."
}

func genDestroyer(name string) string {
	return fmt.Sprintf(
		"%s and %s %s %s",
		genInsult(),
		randyWithArticle(adj),
		randy(compoundAdj),
		genFin(name),
	)
}

// to help make some of the finishers more personal a bit
// of extra effort is needed to get the grammar right.
func genFin(name string) string {
	fin := randy(finisher)

	// handful need an xtra noun thrown in..
	fin = strings.Replace(fin, "<noun>", randy(noun), -1)

	// finishers by default use personal pronouns 
	// so we can leave early if it's personal
	if strings.HasPrefix(name, "You ") {
		return fin
	}

	// if it's not the epic rapunzel finsisher we just
	// need to make sure the pronoun in an oject pronoun
	if !strings.Contains(fin, "Rapunzel") {
		return strings.Replace(fin, "you", "them", -1)
	}

	// get proper fucked with subject+possive pronouns
	fin = strings.Replace(fin, "your", "their", -1)
	return strings.Replace(fin, "you", "they", -1)
}

// ----------------------------------------------------------------------

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
	return &CN{sep: "", clean: false}
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
// Thus, the adj's cleanliness shall be determined at random
// unless the user has admitted they are pussy.
func genCodename(isAPussy bool) (string, string) {
	if isAPussy {
		return randy(cleanAdj), randy(cleanNoun)
	}

	var (
		a = randy(adj)
		n = randy(noun)
	)

	if randomInt(2) == 1 {
		return randy(cleanAdj), n
	}

	return a, n
}