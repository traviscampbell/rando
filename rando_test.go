package rando

import (
	"regexp"
	"strings"
	"testing"
	"unicode"
)

func TestInsult(t *testing.T) {
	wantRE := regexp.MustCompile(`^(?:You are|Rando is)[\w\s]+\.$`)

	// test no name provided
	if got := Insult(""); !wantRE.MatchString(got) {
		t.Errorf("Wanted insult to be prefixed with 'You are', but got %s\n", got)
		t.Fail()
	}

	// test with a name
	if got := Insult("rando"); !wantRE.MatchString(got) {
		t.Errorf("Wanted insult to start with 'Rando', but got %s\n", got)
		t.Fail()
	}

	// test some dickhead trying to get fresh with me
	if got := Insult("travis"); got != youNoFuckMe {
		t.Errorf("Wanted rando to have my back, but instead %s\n", got)
		t.Fail()
	}
}

func TestDestroy(t *testing.T) {
	wantRE := regexp.MustCompile(`^(?:You are|Rando is).*?\.$`)

	// test no name provided
	if got := Destroy(""); !wantRE.MatchString(got) {
		t.Errorf("Wanted destroy to be prefixed with 'You are', bot got %s\n", got)
		t.Fail()
	}

	// test with a name
	if got := Destroy("rando"); !wantRE.MatchString(got) {
		t.Errorf("Wanted destroy to start with 'Rando', but got %s\n", got)
		t.Fail()
	}

	// test some dickhead trying to fresh with me
	if got := Destroy("travis"); got != youNoFuckMe {
		t.Errorf("Wanted rando to have my back, but instead %s\n", got)
		t.Fail()
	}
}

func TestCodename(t *testing.T) {
	if got := NewCodename().Generate(); got == "" {
		t.Error("Wanted a codename, but didn't get shit.")
		t.Fail()
	}

	if got := NewCodename().Pussify().Generate(); got == "" {
		t.Error("Wanted a wholesome codename, but don't even know what that is")
		t.Fail()
	}

	// test codename with a fixed first word
	if got := NewCodename().WithFWord("first").Generate(); !strings.HasPrefix(got, "First") {
		t.Errorf("Got %s, but codename should have started with 'First'\n", got)
		t.Fail()
	}

	// test codename with a fixed last word
	if got := NewCodename().WithLWord("last").Generate(); !strings.HasSuffix(got, "Last") {
		t.Errorf("Got %s, but codename should have ended with 'Last'\n", got)
		t.Fail()
	}

	// test codename with a space as the seperator between the words
	if got := NewCodename().WithSep(" ").Generate(); len(strings.Fields(got)) != 2 {
		t.Errorf("Got %s, but just want the fucking the adj and noun to have a space in da middle\n", got)
		t.Fail()
	}

	// test codename with all CAPS
	if got := NewCodename().WithCAPS().Generate(); !checkUP(got) {
		t.Errorf("GOT %s, BUTT WANTED ALL FUCKING CAPS!!!!", got)
		t.Fail()
	}

	// test accidental complimenting which has a 25% chance of a 50% chance
	// of a 75% chance of a 1% chancea 50% chance of a 1% chance of happening.
	// If run 31337 and rando is still being a dick, probably stay away from casino's.
	flattered := false
	for i := 0; i < 31337; i++ {
		got := Insult("")
		for _, c := range compliment {
			if strings.Contains(got, c) {
				flattered = true
				break
			}
		}
	}

	if !flattered {
		t.Errorf("Wanted to be flattered, accidently or otherwise, but just feel like maybe rando is right...")
		t.Fail()
	}
}

func checkUP(codename string) bool {
	for _, r := range codename {
		if !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}
