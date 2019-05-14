package rando

import (
	"regexp"
	"strings"
	"testing"
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
	if got := NewCodenamer().Please(); got == "" {
		t.Error("Wanted a codename, but didn't get shit.")
		t.Fail()
	}

	if got := NewCodenamer().WhatAPussy().Please(); got == "" {
		t.Error("Wanted a wholesome codename, but don't even know what that is")
		t.Fail()
	}

	// test codename with a fixed first word
	if got := NewCodenamer().WithFirstWord("first").Please(); !strings.HasPrefix(got, "First") {
		t.Errorf("Got %s, but codename should have started with 'First'\n", got)
		t.Fail()
	}

	// test codename with a fixed last word
	if got := NewCodenamer().WithLastWord("last").Please(); !strings.HasSuffix(got, "Last") {
		t.Errorf("Got %s, but codename should have ended with 'Last'\n", got)
		t.Fail()
	}

	// test codename with a space as the seperator between the words
	if got := NewCodenamer().WithSeperator(" ").Please(); len(strings.Fields(got)) != 2 {
		t.Errorf("Got %s, but just want the fucking the adj and noun to have a space in da middle\n", got)
		t.Fail()
	}
}
