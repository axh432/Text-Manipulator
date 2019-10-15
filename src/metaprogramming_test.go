package main

import (
	"errors"
	"github.com/golang-collections/collections/stack"
	"reflect"
)

func iterate(collection []reflect.Value){}

func assign(value reflect.Value, variable reflect.Value){}

func dotAccess(path string, object reflect.Value){}

func callFunction(args []reflect.Value){}

func altFindFirstCodeBlock(s *Section, blockPattern *OpenClosedPattern) (Section, error) {

	matches := s.findAll(blockPattern.whole)
	blockStack := stack.New()

	for _, match := range matches {

		symbol := match.toString()

		if blockPattern.isOpen(symbol) {

			blockStack.Push(match)

		} else if blockPattern.isClosed(symbol) {

			stackLength := blockStack.Len()

			if stackLength > 1 {
				blockStack.Pop()
			} else if stackLength == 1 {
				pop := blockStack.Pop()

				startMatch, ok := pop.(Section)
				if ok {
					return Section{startMatch.start, match.end, startMatch.source}, nil
				}

			}

		}

	}

	return Section{}, errors.New("no matches found in this section")

}
