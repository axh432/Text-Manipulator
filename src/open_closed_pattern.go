package main

import (
	. "regexp"
)

type OpenClosedPattern struct {
	whole  *Regexp
	open   *Regexp
	closed *Regexp
}

func newOpenClosedPatternString(open string, closed string) *OpenClosedPattern {

	openPattern := MustCompile(open)
	closedPattern := MustCompile(closed)
	wholePattern := MustCompile("[" + open + closed + "]")

	return &OpenClosedPattern{wholePattern, openPattern, closedPattern }
}

func newOpenClosedPatternRegex(whole *Regexp, open *Regexp, closed *Regexp) *OpenClosedPattern {
	return &OpenClosedPattern{whole, open, closed }
}

func (codeBlockPattern *OpenClosedPattern) isOpen(symbol string) bool {
	return codeBlockPattern.open.MatchString(symbol)
}

func (codeBlockPattern *OpenClosedPattern) isClosed(symbol string) bool {
	return codeBlockPattern.closed.MatchString(symbol)
}