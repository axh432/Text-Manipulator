package main

import (
	. "regexp"
)

type CodeBlockPattern struct {
	whole  Regexp
	open   Regexp
	closed Regexp
}

func newCodeBlockPattern(whole Regexp, open Regexp, closed Regexp) *CodeBlockPattern{
	return &CodeBlockPattern{ whole, open, closed }
}

func (codeBlockPattern *CodeBlockPattern) isOpen(symbol string) bool {
	return codeBlockPattern.open.MatchString(symbol)
}

func (codeBlockPattern *CodeBlockPattern) isClosed(symbol string) bool {
	return codeBlockPattern.closed.MatchString(symbol)
}