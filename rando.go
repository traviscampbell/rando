package rando

import (
	"crypto/rand" // don't act like you're not impressed...
	"fmt"
	"math/big"
	"regexp"
	"strings"
	"unicode"
)

// randy string from given slice
func randy(source []string) string {
	return source[randomInt(len(source)-1)]
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

	// with no mas titleCase and not wanting a 3rd party library this
	// is what I resorted to incase some dilapidated dickhole wants
	// multiple words title-cased. also, I don't want to talk about it...
	llamo := ""
	for _, word := range strings.Split(name, " ") {
		llamo += string(unicode.ToUpper(rune(word[0]))) + word[1:] + " "
	}
	return llamo + "is"
}

// get a crypto random int between 0 and n (n is not inclusive. e.g.if n=100 it's between 0-99)
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
	nameRE  *regexp.Regexp
	enabled bool
}

func (e *easterEgg) protectMe(n string) bool {
	return e.nameRE.MatchString(n)
}

// SetEasterEggRegexp lets you override the regexp val to help
// protect yourself from getting boiled, beaten, whisked, cracked
// deviled, scrambled, steamed, pickled, poached, baked, or even fried.
func SetEasterEggRegexp(reStr string) {
	humpty.nameRE = regexp.MustCompile(reStr)
}

// PoachEasterEgg disables protective easter egg shell
func PoachEasterEgg() {
	humpty.enabled = false
}

// humpty is the default easter-egg preloaded to defend me because I'm
// a little chicken who can dish it out, but can't take it... ⊙▂⊙
var humpty = easterEgg{
	nameRE:  regexp.MustCompile(`(?i)tcam|tra(vis|cam)`),
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
	if accidentlyCompliment() {
		name = normaleyes(name)
		return fmt.Sprintf("%s %s!", name, genCompliment(name))
	}

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
		strings.ToLower(randy(noun)),
	)
}

// Destroy them on an emotional level with a
// devestating compound insult in the form:
// <Insult> and a(n)? <Adj> <CompoundAdj> <Noun> <Finisher>.
func Destroy(name string) string {
	if accidentlyCompliment() {
		name = normaleyes(name)
		return name + genCompliment(name) + "!"
	}

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
// honestly, I suck at grammar but know a thing or two about
// shit-arounds so bear witness to some those below.:)
func genFin(name string) string {
	fin := randy(finisher)

	noun := strings.ToLower(randy(noun))

	// handful need an xtra noun thrown in..
	fin = strings.Replace(fin, "<noun>", noun, -1)

	// finishers by default use personal pronouns
	// so we can leave early if it's personal
	if strings.HasPrefix(name, "You ") {
		return strings.NewReplacer(
			"<they>", "",
			"<them>", "",
			"<their>", "",
			"<", "",
			">", "",
		).Replace(fin)
	}

	// get proper fucked with subject+possive pronouns
	return strings.NewReplacer(
		"<your>", "",
		"<you>", "",
		"<", "",
		">", "",
	).Replace(fin)
}

// ----------------------------------------------------------------------

// CN represents codename and the various
// options that can be applied.
type CN struct {
	fword     string
	lword     string
	sep       string
	caps      bool
	clean     bool
	halfClean bool
}

// NewCodename returns default pre-configured instance
// used to genrate and configure codenames.
func NewCodename() *CN {
	return &CN{sep: "", clean: false}
}

// WithFWord sets the firstword to the word provided.
func (cn *CN) WithFWord(fw string) *CN {
	cn.fword = fw
	return cn
}

// WithLWord sets the lastword to the word provided.
func (cn *CN) WithLWord(lw string) *CN {
	cn.lword = lw
	return cn
}

// WithSep lets you customize what, if anything, will seperate the parts
func (cn *CN) WithSep(sep string) *CN {
	cn.sep = sep
	return cn
}

// Pussify will make sure ensure the space is kept safe
func (cn *CN) Pussify() *CN {
	cn.clean = true
	return cn
}

// WithCAPS will return the codename in all caps
func (cn *CN) WithCAPS() *CN {
	cn.caps = true
	return cn
}

// HalfClean will pull a random adjective from the clean list but keep
// the noun filthy. It has it's moments...
func (cn *CN) HalfClean() *CN {
	cn.halfClean = true
	return cn
}

// Generate the fucking codename already.
func (cn *CN) Generate() string {
	s1, s2 := cn.generate()
	// shit-around for some words because I like them camelCased more
	// when they are used in the context of a codename.
	s2 = strings.Replace(s2, "-", "", -1)

	if cn.fword != "" {
		s1 = cn.fword
	}

	if cn.lword != "" {
		s2 = cn.lword
	}

	switch {
	case cn.caps:
		s1 = strings.ToUpper(s1)
		s2 = strings.ToUpper(s2)

	default:
		// still don't want to talk about it
		s1 = string(unicode.ToUpper(rune(s1[0]))) + s1[1:]
		s2 = string(unicode.ToUpper(rune(s2[0]))) + s2[1:]
	}

	return s1 + cn.sep + s2
}

// returns a random codename combo. The noun is generally
// the offensive part here and can be enhanced by any adjective.
// If the 'clean' opt is set both adj and noun will be clean(-ish).
// If the 'halfClean' opt is set the adj will at least be from clean list.
// When used properly, the whole shebang should be fairly<->thoroughly offensive.
func (cn *CN) generate() (string, string) {
	if cn.clean {
		return randy(cleanAdj), randy(cleanNoun)
	}

	switch {
	case cn.clean:
		return randy(cleanAdj), randy(cleanNoun)

	case cn.halfClean:
		return randy(cleanAdj), randy(noun)

	default:
		return randy(adj), randy(noun)
	}
}

// ----------------------------------------------------------------------

// there's a 25% chance of a 50% chance of a 75% chance of a 1% chance
// that rando will fuck-up and accidently compliment someone...
// I've not seen it happen in practice, but ultimately accidents do happen.
// I mean just look at you.
func accidentlyCompliment() bool {
	if randomInt(100) < 99 {
		return false
	}

	if randomInt(100) < 75 {
		return false
	}

	if randomInt(100) < 50 {
		return false
	}

	if randomInt(100) < 25 {
		return false
	}

	return true
}

// (ㅅ´ ˘ `)♡ ☜(꒡⌓꒡)
func genCompliment(name string) string {
	if strings.HasPrefix(name, "You ") {
		return strings.NewReplacer(
			"<they>", "",
			"<them>", "",
			"<their>", "",
			"<", "",
			">", "",
		).Replace((randy(compliment)))
	}

	return strings.NewReplacer(
		"<your>", "",
		"<you>", "",
		"<", "",
		">", "",
	).Replace(randy(compliment))
}
