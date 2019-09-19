package main

import (
	. "regexp"
)


type GrammaticalExpression interface {
	GetType() string
}

type Composition struct {
	first interface{}
	second interface{}
}

type Or struct {
	first interface{}
	second interface{}
}

type SinglePattern struct {
	pattern *Regexp
}

func (c Composition) GetType() string {
	return "Composition"
}

func (or Or) GetType() string {
	return "Or"
}

func (sp SinglePattern) GetType() string {
	return "SinglePattern"
}