package new_regex

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

/* These are the regex tutorials found at: https://regexone.com/ implemented in new-regex */

func TestRegexOne(t *testing.T) {

	t.Run("Lesson 1: the ABCs", func(t *testing.T) {
		//pattern that matches the following strings.

		exp := Range(SetOfCharacters("abcdefg"), 1, 7)

		require.Equal(t, true, Match("abcdefg", exp).IsValid)
		require.Equal(t, true, Match("abcde", exp).IsValid)
		require.Equal(t, true, Match("abc", exp).IsValid)
	})

	t.Run("Lesson 1.5: the 123s", func(t *testing.T) {
		//write a pattern that matches only the numbers.
		integer := Label(Range(Number, 1, -1), "integer")
		notANumber := Set(Whitespace, Letter, Punctuation, Symbol)
		notInteger := Range(notANumber, 1, -1)

		numbers := []string{}
		visitor := func(mt *MatchTree) {
			if mt.Label != "" {
				numbers = append(numbers, mt.Value)
			}
		}

		exp := Range(Set(notInteger, integer), 0, -1)

		iter := CreateIterator("var g = 123;")

		result := MatchIter(&iter, exp)

		result.acceptVisitor(visitor)

		fmt.Printf("%v", numbers)

		/*
		abc123xyz
		define "123"
		var g = 123;
		*/

	})

	t.Run("Lesson 2: the 'any' character", func(t *testing.T) {
		//write a pattern the first three but not the last
		/*
			Match	cat.
			Match	896.
			Match	?=+.
			Skip	abc1
		*/
	})

	t.Run("Lesson 3: Matching specific characters", func(t *testing.T) {
		/*
		Match	can	To be completed
		Match	man	To be completed
		Match	fan	To be completed
		Skip	dan	To be completed
		Skip	ran	To be completed
		Skip	pan
		*/
	})

	t.Run("Lesson 4: Excluding specific characters", func(t *testing.T) {
		/*
			Match	can	To be completed
			Match	man	To be completed
			Match	fan	To be completed
			Skip	dan	To be completed
			Skip	ran	To be completed
			Skip	pan
		*/
	})

	//lesson 5 character ranges is not supported atm
	t.Run("Lesson 5: Character ranges", func(t *testing.T) {
		/*
		Match	Ana	To be completed
		Match	Bob	To be completed
		Match	Cpc	To be completed
		Skip	aax	To be completed
		Skip	bby	To be completed
		Skip	ccz
		*/
	})

	t.Run("Lesson 6: Catching some zzz's", func(t *testing.T) {
		//use ranges to match the strings that need to be matched.
		/*
		Match	wazzzzzup	To be completed
		Match	wazzzup	To be completed
		Skip	wazup
		*/
	})

	t.Run("Lesson 7: Matching Repeated Characters", func(t *testing.T) {
		/*
		Match	aaaabcc	To be completed
		Match	aabbbbc	To be completed
		Match	aacc	To be completed
		Skip	a
		*/
	})

	t.Run("Lesson 8: Characters optional", func(t *testing.T) {
		/*
		Match	1 file found?	To be completed
		Match	2 files found?	To be completed
		Match	24 files found?	To be completed
		Skip	No files found.
		*/
	})

	t.Run("Lesson 9: All this whitespace", func(t *testing.T) {
		/*
			Match	1.   abc	To be completed
			Match	2.	abc	To be completed
			Match	3.           abc	To be completed
			Skip	4.abc
		*/
	})

	t.Run("Lesson 10: Starting and ending", func(t *testing.T) {
		/*
			Match	Mission: successful	To be completed
			Skip	Last Mission: unsuccessful	To be completed
			Skip	Next Mission: successful upon capture of target
		*/
	})

	t.Run("Lesson 11: Match groups", func(t *testing.T) {
		//capture only the file name and not the extension
		/*
			Capture	file_record_transcript.pdf	file_record_transcript	To be completed
			Capture	file_07241999.pdf	file_07241999	To be completed
			Skip	testfile_fake.pdf.tmp
		*/
	})

	t.Run("Lesson 12: Nested groups", func(t *testing.T) {
		//capture the full date and the year of the date
		/*
		Capture	Jan 1987	Jan 1987 1987	To be completed
		Capture	May 1969	May 1969 1969	To be completed
		Capture	Aug 2011	Aug 2011 2011
		*/
	})

	t.Run("Lesson 13: More group work", func(t *testing.T) {
		//capture the individual dimensions
		/*
		Capture	1280x720	1280 720	To be completed
		Capture	1920x1600	1920 1600	To be completed
		Capture	1024x768	1024 768
		*/
	})

	t.Run("Lesson 14: It's all conditional", func(t *testing.T) {
		/*
			Match	I love cats	To be completed
			Match	I love dogs	To be completed
			Skip	I love logs	To be completed
			Skip	I love cogs
		*/
	})

	t.Run("Lesson 15: Other special characters", func(t *testing.T) {
		/*
			Match	The quick brown fox jumps over the lazy dog.	To be completed
			Match	There were 614 instances of students getting 90.0% or above.	To be completed
			Match	The FCC had to censor the network for saying &$#*@!.
		*/
	})

}