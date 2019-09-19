package main

import . "regexp"

type MatchedSinglePattern struct {
	Symbol *Regexp
	MitigatingSymbols []string
}

func (m MatchedSinglePattern) GetType() string {
	return "MatchedSinglePattern"
}

