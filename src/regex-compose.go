package main

import (

	. "regexp"
)

func splitKeepDelimiter(pattern *Regexp, toBeSplit string) []string {

	splitStrings := []string{}

	matches := pattern.FindAllStringIndex(toBeSplit, -1)

	if matches != nil {

		for i, match := range matches {

			matchStart := match[0]
			matchEnd := match[1]

			if i == 0 {

				if matchStart > 0 {
					splitStrings = append(splitStrings, toBeSplit[:matchStart])
				}

			}

			if i > 0 {

				previousMatch := matches[i-1]
				previousMatchEnd := previousMatch[1]

				if matchStart - previousMatchEnd > 0 {
					splitStrings = append(splitStrings, toBeSplit[previousMatchEnd:matchStart])
				}

			}

			splitStrings = append(splitStrings, toBeSplit[matchStart:matchEnd])

			if i == len(matches) - 1 {
				splitStrings = append(splitStrings, toBeSplit[matchEnd:])
			}

		}

	}


	return splitStrings
}